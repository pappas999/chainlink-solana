package smoke

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"

	// "github.com/gagliardetto/solana-go"
	// utils2 "github.com/smartcontractkit/chainlink-solana/tests/e2e/utils"

	// relayUtils "github.com/smartcontractkit/chainlink-relay/ops/utils"
	"github.com/smartcontractkit/chainlink-solana/tests/e2e/solclient"
	g "github.com/smartcontractkit/chainlink-solana/tests/e2e/utils"
	"github.com/smartcontractkit/helmenv/environment"
	"github.com/smartcontractkit/helmenv/tools"
	"github.com/smartcontractkit/integrations-framework/actions"
	"github.com/smartcontractkit/integrations-framework/client"
)

type Deployer struct {
	gauntlet g.Gauntlet
	network  string
	Account  map[int]string
}

var _ = Describe("Gauntlet Testing @gauntlet", func() {
	var (
		e        *environment.Environment
		gd       Deployer
		gauntlet g.Gauntlet
		nkb      []NodeKeysBundle
		nets     *client.Networks
		err      error
	)

	solanaCommandError := g.ExecError{
		WhatIsIt:     "Solana Command execution error",
		HowToRespond: " ",
	}

	BeforeEach(func() {
		By("Deploying the environment", func() {
			e, err = environment.DeployOrLoadEnvironment(
				solclient.NewSolanaValidator(),
				tools.ChartsRoot,
			)
			Expect(err).ShouldNot(HaveOccurred())
			err = e.ConnectAll()
			Expect(err).ShouldNot(HaveOccurred())
			err = UploadProgramBinaries(e)
			Expect(err).ShouldNot(HaveOccurred())
		})
		By("Getting the clients", func() {
			networkRegistry := client.NewNetworkRegistry()
			networkRegistry.RegisterNetwork(
				"solana",
				solclient.ClientInitFunc(),
				solclient.ClientURLSFunc(),
			)
			nets, err = networkRegistry.GetNetworks(e)
			Expect(err).ShouldNot(HaveOccurred())
		})
		By("Setup Gauntlet", func() {
			_, err := exec.LookPath("yarn")
			Expect(err).ShouldNot(HaveOccurred())

			// skip all teh gauntlet prompts since trying to handle them gracefully would be very painful
			os.Setenv("SKIP_PROMPTS", "true")
			// make the gauntlet solana calls timeout longer for tests, it normally defaults to 60 seconds
			os.Setenv("CONFIRM_TX_TIMEOUT_SECONDS", "120")

			log.Debug().Str("OS", runtime.GOOS).Msg("Runtime OS:")
			version := "linux"
			if runtime.GOOS == "darwin" {
				version = "macos"
			}

			// Check gauntlet works
			cwd, _ := os.Getwd()
			err = os.Chdir(filepath.Join(cwd + "../../../../gauntlet"))
			Expect(err).ShouldNot(HaveOccurred())

			gauntletBin := filepath.Join(cwd, "../../../gauntlet/bin/gauntlet-") + version
			gauntlet, err = g.NewGauntlet(gauntletBin)
			Expect(err).ShouldNot(HaveOccurred())

			gd = Deployer{
				gauntlet: gauntlet,
				network:  "local",
				Account:  make(map[int]string),
			}
		})
		By("Fund Wallets", func() {
			err = FundOracles(nets.Default, nkb, big.NewFloat(5e4))
			Expect(err).ShouldNot(HaveOccurred())

			// TODO: create a proper waiter to do this
			log.Debug().Msg("Sleeping to let the wallets fill")
			time.Sleep(10 * time.Second)
		})
	})

	Describe("gauntlet commands", func() {

		It("token", func() {
			network := "blarg"
			networkConfig, err := GetDefaultGauntletConfig(network, e)
			Expect(err).ShouldNot(HaveOccurred())

			// Deploy Link
			log.Debug().Msg("Deploying LINK Token...")
			args := []string{
				"token:deploy",
				gd.gauntlet.Flag("network", network),
			}

			errHandling := []g.ExecError{
				solanaCommandError,
			}
			report, err := gd.gauntlet.ExecuteAndRead(args, errHandling)
			Expect(err).ShouldNot(HaveOccurred())

			linkAddress := report.Responses[0].Contract
			networkConfig["LINK"] = linkAddress

			// Create Billing and Requester Access Controllers
			log.Debug().Msg("Deploying Access Controller for Requester...")
			acArgs := []string{
				"access_controller:initialize",
				gd.gauntlet.Flag("network", network),
				linkAddress,
			}

			acErrHandling := []g.ExecError{
				solanaCommandError,
			}
			report, err = gd.gauntlet.ExecuteAndRead(acArgs, acErrHandling)
			Expect(err).ShouldNot(HaveOccurred())
			requesterAccessController := report.Responses[0].Contract

			log.Debug().Msg("Deploying Access Controller for Billing...")
			report, err = gd.gauntlet.ExecuteAndRead(acArgs, acErrHandling)
			Expect(err).ShouldNot(HaveOccurred())
			billingAccessController := report.Responses[0].Contract

			networkConfig["REQUESTER_ACCESS_CONTROLLER"] = requesterAccessController
			networkConfig["BILLING_ACCESS_CONTROLLER"] = billingAccessController
			err = WriteNetworkConfigMap(fmt.Sprintf("networks/.env.%s", network), networkConfig)
			Expect(err).ShouldNot(HaveOccurred())

			// Create Store
			log.Debug().Msg("Deploying Store...")
			storeArgs := []string{
				"store:initialize",
				gd.gauntlet.Flag("network", network),
			}

			storeErrHandling := []g.ExecError{
				solanaCommandError,
			}
			report, err = gd.gauntlet.ExecuteAndRead(storeArgs, storeErrHandling)
			Expect(err).ShouldNot(HaveOccurred())

			storeAccount := report.Responses[0].Contract

			// Create store feed
			input := map[string]interface{}{
				"store":       storeAccount,
				"granularity": 30,
				"liveLength":  1024,
				"decimals":    8,
				"description": "Test LINK/USD",
			}
			jsonInput, err := json.Marshal(input)
			Expect(err).ShouldNot(HaveOccurred())

			storeCreateFeedArgs := []string{
				"store:create_feed",
				gd.gauntlet.Flag("network", network),
				gd.gauntlet.Flag("state", storeAccount),
				gd.gauntlet.Flag("input", string(jsonInput)),
			}
			report, err = gd.gauntlet.ExecuteAndRead(storeCreateFeedArgs, storeErrHandling)
			Expect(err).ShouldNot(HaveOccurred())

			// log.Debug().Msg("Deploying OCR2...")
			// ocr2Args := []string{
			// 	"ocr2:initialize",
			// 	gd.gauntlet.Flag("network", network),
			// }

			// ocr2ErrHandling := []g.ExecError{
			// 	solanaCommandError,
			// }
			// _, err = gd.gauntlet.ExecCommand(ocr2Args, ocr2ErrHandling)
			// // if we got an error we can check to see if it just didn't finish in 60 seconds by parsing the output or error for the tx signature
			// Expect(err).ShouldNot(HaveOccurred())

			// report, err = gd.gauntlet.ReadCommandReport()
			// Expect(err).ShouldNot(HaveOccurred())

			// token:deploy

			// token:read_state

			// token:transfer
		})
		// It("access_controller", func() {
		// 	Expect("abc").To(Equal("abc"))
		// 	// access_controller:initialize

		// 	// access_controller:add_access

		// 	// access_controller:read_state
		// })
		// It("ocr2", func() {
		// 	Expect("abc").To(Equal("abc"))
		// 	// ocr2:initialize
		// 	// ocr2:initialize:flow
		// 	// ocr2:set_billing
		// 	// ocr2:pay_remaining
		// 	// ocr2:set_payees
		// 	// ocr2:set_config
		// 	// ocr2:set_validator_config
		// 	// ocr2:read_state
		// 	// ocr2:set_offchain_config:flow
		// 	// ocr2:begin_offchain_config
		// 	// ocr2:write_offchain_config
		// 	// ocr2:commit_offchain_config
		// 	// ocr2:inspect
		// 	// ocr2:transmit
		// 	// ocr2:setup:flow
		// 	// ocr2:setup:rdd:flow
		// })
		// It("deviation_flagging_validator", func() {
		// 	Expect("abc").To(Equal("abc"))
		// 	// deviation_flagging_validator:initialize
		// })
	})

	AfterEach(func() {
		By("Tearing down the environment", func() {
			err = actions.TeardownSuite(e, nil, "logs")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
