package smoke

import (
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
		e              *environment.Environment
		chainlinkNodes []client.Chainlink
		// cd             contracts.ContractDeployer
		gd       Deployer
		gauntlet g.Gauntlet
		// lt             contracts.LinkToken
		// validator      contracts.OCRv2DeviationFlaggingValidator
		// billingAC      contracts.OCRv2AccessController
		// requesterAC    contracts.OCRv2AccessController
		// ocr2 contracts.OCRv2
		// ocConfig   contracts.OffChainAggregatorV2Config
		nkb        []NodeKeysBundle
		mockserver *client.MockserverClient
		nets       *client.Networks
		err        error
	)

	solanaCommandError := g.ExecError{
		WhatIsIt:     "Solana Command execution error",
		HowToRespond: " ",
	}

	BeforeEach(func() {
		By("Deploying the environment", func() {
			e, err = environment.DeployOrLoadEnvironment(
				solclient.NewChainlinkSolOCRv2(),
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
			// nets.Default.GetClients()[0]
			// nets.Default.
			mockserver, err = client.ConnectMockServer(e)
			Expect(err).ShouldNot(HaveOccurred())
			chainlinkNodes, err = client.ConnectChainlinkNodes(e)
			Expect(err).ShouldNot(HaveOccurred())
			_, nkb, err = DefaultOffChainConfigParamsFromNodes(chainlinkNodes)
			// ocConfig, nkb, err = DefaultOffChainConfigParamsFromNodes(chainlinkNodes)
			Expect(err).ShouldNot(HaveOccurred())
		})
		By("Setup Gauntlet", func() {
			_, err := exec.LookPath("yarn")
			Expect(err).ShouldNot(HaveOccurred())
			// exec.Command("").
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
		By("Deploying contracts", func() {
			// cd, err = solclient.NewContractDeployer(nets.Default, e)
			// Expect(err).ShouldNot(HaveOccurred())
			// _, err := cd.DeployLinkTokenContract()
			// lt, err = cd.DeployLinkTokenContract()
			// Expect(err).ShouldNot(HaveOccurred())
			// log.Debug().Str("Address", lt.Address()).Msg("Link Token Address")
			err = FundOracles(nets.Default, nkb, big.NewFloat(5e4))
			Expect(err).ShouldNot(HaveOccurred())
			// billingAC, err = cd.DeployOCRv2AccessController()
			// Expect(err).ShouldNot(HaveOccurred())
			// requesterAC, err = cd.DeployOCRv2AccessController()
			// Expect(err).ShouldNot(HaveOccurred())
			// err = nets.Default.WaitForEvents()
			// Expect(err).ShouldNot(HaveOccurred())

			// validator, err = cd.DeployOCRv2DeviationFlaggingValidator(billingAC.Address())
			// Expect(err).ShouldNot(HaveOccurred())
			// ocr2, err = cd.DeployOCRv2(billingAC.Address(), requesterAC.Address(), lt.Address())
			// Expect(err).ShouldNot(HaveOccurred())
			// err = nets.Default.WaitForEvents()
			// Expect(err).ShouldNot(HaveOccurred())

			// err = ocr2.SetBilling(uint32(1), uint32(1), billingAC.Address())
			// Expect(err).ShouldNot(HaveOccurred())
			// validatorAuth, err := ocr2.AuthorityAddr("validator")
			// Expect(err).ShouldNot(HaveOccurred())
			// err = billingAC.AddAccess(validatorAuth)
			// Expect(err).ShouldNot(HaveOccurred())
			// err = ocr2.SetOracles(ocConfig)
			// Expect(err).ShouldNot(HaveOccurred())

			// err = nets.Default.WaitForEvents()
			// Expect(err).ShouldNot(HaveOccurred())

			// err = ocr2.SetOffChainConfig(ocConfig)
			// Expect(err).ShouldNot(HaveOccurred())
			// err = ocr2.DumpState()
			// Expect(err).ShouldNot(HaveOccurred())
		})
		By("Set MockServer up", func() {
			err = mockserver.SetValuePath("/variable", 5)
			Expect(err).ShouldNot(HaveOccurred())
			err = mockserver.SetValuePath("/juels", 1)
			Expect(err).ShouldNot(HaveOccurred())

			// err = CreateOCR2Jobs(
			// 	chainlinkNodes,
			// 	nkb,
			// 	mockserver,
			// 	ocr2,
			// 	validator,
			// )
			// Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("gauntlet commands", func() {

		It("token", func() {
			solUrls, err := e.Charts.Connections("solana-validator").LocalURLsByPort("http-rpc", environment.HTTP)
			Expect(err).ShouldNot(HaveOccurred())
			network := "blarg"
			networkConfig := map[string]string{
				"NETWORK":                      "local",
				"NODE_URL":                     solUrls[0].String(),
				"PROGRAM_ID_OCR2":              "CF13pnKGJ1WJZeEgVAtFdUi4MMndXm9hneiHs8azUaZt",
				"PROGRAM_ID_ACCESS_CONTROLLER": "2F5NEkMnCRkmahEAcQfTQcZv1xtGgrWFfjENtTwHLuKg",
				"PROGRAM_ID_STORE":             "A7Jh2nb1hZHwqEofm4N8SXbKTj82rx7KUfjParQXUyMQ",
				"PRIVATE_KEY":                  "[82,252,248,116,175,84,117,250,95,209,157,226,79,186,119,203,91,102,11,93,237,3,147,113,49,205,35,71,74,208,225,183,24,204,237,135,197,153,100,220,237,111,190,58,211,186,148,129,219,173,188,168,137,129,84,192,188,250,111,167,151,43,111,109]",
				"SECRET":                       "[only,unfair,fiction,favorite,sudden,strategy,rotate,announce,rebuild,keep,violin,nuclear]",
			}

			err = WriteNetworkConfigMap(fmt.Sprintf("networks/.env.%s", network), networkConfig)
			Expect(err).ShouldNot(HaveOccurred())

			// sleep to let the wallet fill
			log.Debug().Msg("Sleeping to let the wallets fill")
			time.Sleep(10 * time.Second)

			log.Debug().Msg("Deploying LINK Token...")
			args := []string{
				"token:deploy",
				gd.gauntlet.Flag("network", network),
			}

			errHandling := []g.ExecError{
				solanaCommandError,
			}
			_, err = gd.gauntlet.ExecCommand(args, errHandling)
			// if we got an error we can check to see if it just didn't finish in 60 seconds by parsing the output or error for the tx signature
			Expect(err).ShouldNot(HaveOccurred())

			report, err := gd.gauntlet.ReadCommandReport()
			Expect(err).ShouldNot(HaveOccurred())

			linkAddress := report.Responses[0].Contract
			networkConfig["LINK"] = linkAddress
			err = WriteNetworkConfigMap(fmt.Sprintf("networks/.env.%s", network), networkConfig)
			Expect(err).ShouldNot(HaveOccurred())

			// log.Debug().Msg("Deploying Access Controller...")
			// acArgs := []string{
			// 	"access_controller:initialize",
			// 	gd.gauntlet.Flag("network", network),
			// }

			// acErrHandling := []g.ExecError{
			// 	solanaCommandError,
			// }
			// _, err = gd.gauntlet.ExecCommand(acArgs, acErrHandling)
			// // if we got an error we can check to see if it just didn't finish in 60 seconds by parsing the output or error for the tx signature
			// Expect(err).ShouldNot(HaveOccurred())

			// report, err := gd.gauntlet.ReadCommandReport()
			// Expect(err).ShouldNot(HaveOccurred())

			// requesterAccessController := report.Responses[0].Contract
			// networkConfig["REQUESTER_ACCESS_CONTROLLER"] = requesterAccessController
			// // networkConfig["BILLING_ACCESS_CONTROLLER"] = requesterAccessController
			// err = WriteNetworkConfigMap(fmt.Sprintf("networks/.env.%s", network), networkConfig)
			// Expect(err).ShouldNot(HaveOccurred())

			// log.Debug().Msg("Deploying Store...")
			// storeArgs := []string{
			// 	"store:initialize",
			// 	gd.gauntlet.Flag("network", network),
			// }

			// storeErrHandling := []g.ExecError{
			// 	solanaCommandError,
			// }
			// _, err = gd.gauntlet.ExecCommand(storeArgs, storeErrHandling)
			// // if we got an error we can check to see if it just didn't finish in 60 seconds by parsing the output or error for the tx signature
			// Expect(err).ShouldNot(HaveOccurred())

			// report, err = gd.gauntlet.ReadCommandReport()
			// Expect(err).ShouldNot(HaveOccurred())

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
