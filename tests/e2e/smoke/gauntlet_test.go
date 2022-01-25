package smoke

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"

	// "github.com/gagliardetto/solana-go"
	// utils2 "github.com/smartcontractkit/chainlink-solana/tests/e2e/utils"

	// relayUtils "github.com/smartcontractkit/chainlink-relay/ops/utils"
	"github.com/smartcontractkit/chainlink-solana/tests/e2e/common"
	"github.com/smartcontractkit/chainlink-solana/tests/e2e/solclient"
	"github.com/smartcontractkit/chainlink-solana/tests/e2e/utils"
	g "github.com/smartcontractkit/chainlink-solana/tests/e2e/utils"
	"github.com/smartcontractkit/helmenv/environment"
	"github.com/smartcontractkit/helmenv/tools"
	"github.com/smartcontractkit/integrations-framework/actions"
	"github.com/smartcontractkit/integrations-framework/contracts"
)

type Deployer struct {
	gauntlet g.Gauntlet
	network  string
	Account  map[int]string
}

const RETRY_COUNT = 3

var _ = Describe("Gauntlet Testing @gauntlet", func() {
	var (
		// e              *environment.Environment
		gd       Deployer
		gauntlet g.Gauntlet
		// chainlinkNodes []client.Chainlink
		cd contracts.ContractDeployer
		// store          contracts.OCRv2Store
		// billingAC      contracts.OCRv2AccessController
		// requesterAC    contracts.OCRv2AccessController
		// ocr2           contracts.OCRv2
		// ocConfig       contracts.OffChainAggregatorV2Config
		// nkb []NodeKeysBundle
		// mockserver     *client.MockserverClient
		// nets *client.Networks
		err error
	)
	var state = &common.OCRv2TestState{}

	solanaCommandError := []string{
		"Solana Command execution error",
	}

	BeforeEach(func() {
		By("Deploying the environment", func() {
			state.Env, err = environment.DeployOrLoadEnvironment(
				solclient.NewChainlinkSolOCRv2(1, false),
				tools.ChartsRoot,
			)
			Expect(err).ShouldNot(HaveOccurred())
			err = state.Env.ConnectAll()
			Expect(err).ShouldNot(HaveOccurred())
			state.UploadProgramBinaries()
		})
		By("Getting the clients", func() {
			// networkRegistry := client.NewNetworkRegistry()
			// networkRegistry.RegisterNetwork(
			// 	"solana",
			// 	solclient.ClientInitFunc(),
			// 	solclient.ClientURLSFunc(),
			// )
			// state.Networks, err = networkRegistry.GetNetworks(state.Env)
			// Expect(err).ShouldNot(HaveOccurred())
			// state.MockServer, err = client.ConnectMockServer(state.Env)
			// Expect(err).ShouldNot(HaveOccurred())
			// state.ChainlinkNodes, err = client.ConnectChainlinkNodes(state.Env)
			// Expect(err).ShouldNot(HaveOccurred())

			state.SetupClients()
			state.OffChainConfig, state.NodeKeysBundle, err = common.DefaultOffChainConfigParamsFromNodes(state.ChainlinkNodes)
			Expect(err).ShouldNot(HaveOccurred())
			cd, err = solclient.NewContractDeployer(state.Networks.Default, state.Env)
			Expect(err).ShouldNot(HaveOccurred())
		})
		By("Setup Gauntlet", func() {
			_, err := exec.LookPath("yarn")
			Expect(err).ShouldNot(HaveOccurred())

			// skip all teh gauntlet prompts since trying to handle them gracefully would be very painful
			os.Setenv("SKIP_PROMPTS", "true")
			// make the gauntlet solana calls timeout longer for tests, it normally defaults to 60 seconds
			os.Setenv("CONFIRM_TX_TIMEOUT_SECONDS", "15")
			os.Setenv("CONFIRM_TX_COMMITMENT", "confirmed")

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
			err = common.FundOracles(state.Networks.Default, state.NodeKeysBundle, big.NewFloat(5e4))
			Expect(err).ShouldNot(HaveOccurred())
			err = state.Networks.Default.(*solclient.Client).WaitForEvents()
			Expect(err).ShouldNot(HaveOccurred())

			// TODO: create a proper waiter to do this
			// log.Debug().Msg("Sleeping to let the wallets fill")
			// time.Sleep(30 * time.Second)
		})
	})

	Describe("gauntlet commands", func() {

		XIt("token", func() {
			network := "blarg"
			networkConfig, err := utils.GetDefaultGauntletConfig(network, state.Env)
			Expect(err).ShouldNot(HaveOccurred())

			// Deploy Link
			log.Debug().Msg("Deploying LINK Token...")
			args := []string{
				"token:deploy",
				gd.gauntlet.Flag("network", network),
			}

			report, output, err := gd.gauntlet.ExecuteAndRead(args, solanaCommandError, RETRY_COUNT)
			log.Info().Str("Out", output).Msg("Gauntlet Output")
			Expect(err).ShouldNot(HaveOccurred())
			if err != nil {
				// we ran into an error, grab the signature and wait for the tx to complete
				r, _ := regexp.Compile("Check signature ([a-zA-Z0-9]+) using the Solana Explorer")
				matches := r.FindStringSubmatch(output)
				Expect(len(matches)).To(Equal(2))
				fmt.Println(matches[1])
				sig, err := solana.SignatureFromBase58(matches[1])
				Expect(err).ShouldNot(HaveOccurred())
				state.Networks.Default.(*solclient.Client).QueueTX(sig, rpc.CommitmentFinalized)
				err = state.Networks.Default.(*solclient.Client).WaitForEvents()
				Expect(err).ShouldNot(HaveOccurred())
			}
			// lt, err := cd.DeployLinkTokenContract()
			// Expect(err).ShouldNot(HaveOccurred())
			// linkAddress := lt.Address()
			linkAddress := report.Responses[0].Contract
			networkConfig["LINK"] = linkAddress
			err = utils.WriteNetworkConfigMap(fmt.Sprintf("networks/.env.%s", network), networkConfig)
			Expect(err).ShouldNot(HaveOccurred())

			// Read the token state
			log.Debug().Msg("Read the state of the token.")
			acArgs := []string{
				"token:read_state",
				gd.gauntlet.Flag("network", network),
				linkAddress,
			}

			_, _, err = gd.gauntlet.ExecuteAndRead(acArgs, solanaCommandError, RETRY_COUNT)
			Expect(err).ShouldNot(HaveOccurred())
		})

		XIt("token2", func() {
			network := "blarg"
			networkConfig, err := utils.GetDefaultGauntletConfig(network, state.Env)
			Expect(err).ShouldNot(HaveOccurred())

			// Deploy Link
			log.Debug().Msg("Deploying LINK Token...")
			lt, err := cd.DeployLinkTokenContract()
			Expect(err).ShouldNot(HaveOccurred(), "Deploying token failed")
			err = state.Networks.Default.(*solclient.Client).WaitForEvents()
			Expect(err).ShouldNot(HaveOccurred())
			linkAddress := lt.Address()
			networkConfig["LINK"] = linkAddress
			err = utils.WriteNetworkConfigMap(fmt.Sprintf("networks/.env.%s", network), networkConfig)
			Expect(err).ShouldNot(HaveOccurred())

			// Read the token state
			log.Debug().Msg("Read the state of the token.")
			args := []string{
				"token:read_state",
				gd.gauntlet.Flag("network", network),
				linkAddress,
			}

			_, output, err := gd.gauntlet.ExecuteAndRead(args, solanaCommandError, RETRY_COUNT)
			Expect(err).ShouldNot(HaveOccurred(), "Reading the token state failed")
			Expect(strings.Contains(output, "supply: <BN: de0b6b3a7640000>")).To(Equal(true), "We should have the expected supply of tokens in the deployed address")

			// token:transfer
			log.Debug().Msg("Transfer token.")
			args = []string{
				"token:transfer",
				gd.gauntlet.Flag("network", network),
				"--to=7xBSFPrRhXdZW3BmJpa5tydtFngDhapnh8SzihtFKd2U",
				"--amount=100",
				linkAddress,
			}

			_, _, err = gd.gauntlet.ExecuteAndRead(args, solanaCommandError, RETRY_COUNT)
			Expect(err).ShouldNot(HaveOccurred())
		})

		// It("access_controller", func() {
		// 	network := "blarg"
		// 	networkConfig, err := GetDefaultGauntletConfig(network, e)
		// 	Expect(err).ShouldNot(HaveOccurred())

		// 	// Deploy Link
		// 	log.Debug().Msg("Deploying LINK Token...")
		// 	args := []string{
		// 		"token:deploy",
		// 		gd.gauntlet.Flag("network", network),
		// 	}

		// 	errHandling := []g.ExecError{
		// 		solanaCommandError,
		// 	}
		// 	report, err := gd.gauntlet.ExecuteAndRead(args, errHandling)
		// 	Expect(err).ShouldNot(HaveOccurred())

		// 	linkAddress := report.Responses[0].Contract
		// 	networkConfig["LINK"] = linkAddress

		// 	// Create Billing and Requester Access Controllers
		// 	log.Debug().Msg("Deploying Access Controller for Requester...")
		// 	acArgs := []string{
		// 		"access_controller:initialize",
		// 		gd.gauntlet.Flag("network", network),
		// 		linkAddress,
		// 	}

		// 	acErrHandling := []g.ExecError{
		// 		solanaCommandError,
		// 	}
		// 	report, err = gd.gauntlet.ExecuteAndRead(acArgs, acErrHandling)
		// 	Expect(err).ShouldNot(HaveOccurred())
		// 	requesterAccessController := report.Responses[0].Contract

		// 	log.Debug().Msg("Deploying Access Controller for Billing...")
		// 	report, err = gd.gauntlet.ExecuteAndRead(acArgs, acErrHandling)
		// 	Expect(err).ShouldNot(HaveOccurred())
		// 	billingAccessController := report.Responses[0].Contract

		// 	networkConfig["REQUESTER_ACCESS_CONTROLLER"] = requesterAccessController
		// 	networkConfig["BILLING_ACCESS_CONTROLLER"] = billingAccessController
		// 	err = WriteNetworkConfigMap(fmt.Sprintf("networks/.env.%s", network), networkConfig)
		// 	Expect(err).ShouldNot(HaveOccurred())

		// 	// Create Store
		// 	log.Debug().Msg("Deploying Store...")
		// 	storeArgs := []string{
		// 		"store:initialize",
		// 		gd.gauntlet.Flag("network", network),
		// 	}

		// 	storeErrHandling := []g.ExecError{
		// 		solanaCommandError,
		// 	}
		// 	report, err = gd.gauntlet.ExecuteAndRead(storeArgs, storeErrHandling)
		// 	Expect(err).ShouldNot(HaveOccurred())

		// 	storeAccount := report.Responses[0].Contract

		// 	// Create store feed
		// 	input := map[string]interface{}{
		// 		"store":       storeAccount,
		// 		"granularity": 30,
		// 		"liveLength":  1024,
		// 		"decimals":    8,
		// 		"description": "Test LINK/USD",
		// 	}
		// 	jsonInput, err := json.Marshal(input)
		// 	Expect(err).ShouldNot(HaveOccurred())

		// 	storeCreateFeedArgs := []string{
		// 		"store:create_feed",
		// 		gd.gauntlet.Flag("network", network),
		// 		gd.gauntlet.Flag("state", storeAccount),
		// 		gd.gauntlet.Flag("input", string(jsonInput)),
		// 	}
		// 	report, err = gd.gauntlet.ExecuteAndRead(storeCreateFeedArgs, storeErrHandling)
		// 	Expect(err).ShouldNot(HaveOccurred())

		// 	// log.Debug().Msg("Deploying OCR2...")
		// 	// ocr2Args := []string{
		// 	// 	"ocr2:initialize",
		// 	// 	gd.gauntlet.Flag("network", network),
		// 	// }

		// 	// ocr2ErrHandling := []g.ExecError{
		// 	// 	solanaCommandError,
		// 	// }
		// 	// _, err = gd.gauntlet.ExecCommand(ocr2Args, ocr2ErrHandling)
		// 	// // if we got an error we can check to see if it just didn't finish in 60 seconds by parsing the output or error for the tx signature
		// 	// Expect(err).ShouldNot(HaveOccurred())

		// 	// report, err = gd.gauntlet.ReadCommandReport()
		// 	// Expect(err).ShouldNot(HaveOccurred())

		// 	// access_controller:initialize

		// 	// access_controller:add_access

		// 	// access_controller:read_state
		// })
		It("deploy ocr2", func() {
			network := "deployocr"
			networkConfig, err := utils.GetDefaultGauntletConfig(network, state.Env)
			Expect(err).ShouldNot(HaveOccurred(), "Writing the gauntlet config failed")

			// Deploy Link
			log.Debug().Msg("Deploying LINK Token...")
			lt, err := cd.DeployLinkTokenContract()
			Expect(err).ShouldNot(HaveOccurred(), "Deploying token failed")
			err = state.Networks.Default.(*solclient.Client).WaitForEvents()
			Expect(err).ShouldNot(HaveOccurred())
			linkAddress := lt.Address()
			networkConfig["LINK"] = linkAddress
			// ocr2:initialize
			// ocr2:initialize:flow
			// ocr2:set_billing
			// ocr2:pay_remaining
			// ocr2:set_payees
			// ocr2:set_config
			// ocr2:set_validator_config
			// ocr2:read_state
			// ocr2:set_offchain_config:flow
			// ocr2:begin_offchain_config
			// ocr2:write_offchain_config
			// ocr2:commit_offchain_config
			// ocr2:inspect
			// ocr2:transmit
			// ocr2:setup:flow
			// ocr2:setup:rdd:flow
			log.Info().Msg("Deploying OCR Feed")
			log.Info().Msg("Init Requester Access Controller")
			accessControllerArgs := []string{
				"access_controller:initialize",
				gd.gauntlet.Flag("network", network),
			}
			report, _, err := gd.gauntlet.ExecuteAndRead(
				accessControllerArgs,
				solanaCommandError,
				RETRY_COUNT,
			)
			Expect(err).ShouldNot(HaveOccurred(), "Requester access controller initialization failure")
			RequesterAccessController := report.Responses[0].Contract

			log.Info().Msg("Init Billing Access Controller")
			report, _, err = gd.gauntlet.ExecuteAndRead(
				accessControllerArgs,
				solanaCommandError,
				RETRY_COUNT,
			)
			Expect(err).ShouldNot(HaveOccurred(), "Billing access controller initialization failure")
			BillingAccessController := report.Responses[0].Contract

			log.Info().Msg("Init Store")
			report, _, err = gd.gauntlet.ExecuteAndRead(
				[]string{"store:initialize",
					gd.gauntlet.Flag("network", network),
					gd.gauntlet.Flag("accessController", BillingAccessController)},
				solanaCommandError,
				RETRY_COUNT,
			)
			Expect(err).ShouldNot(HaveOccurred(), "Store initialization failed")
			StoreAccount := report.Responses[0].Contract

			log.Info().Msg("Create Feed in Store")
			input := map[string]interface{}{
				"store":       StoreAccount,
				"granularity": 30,
				"liveLength":  1024,
				"decimals":    8,
				"description": "Test LINK/USD",
			}

			jsonInput, err := json.Marshal(input)
			Expect(err).ShouldNot(HaveOccurred(), "Marshaling the stores feed input failed")

			report, _, err = gd.gauntlet.ExecuteAndRead(
				[]string{"store:create_feed",
					gd.gauntlet.Flag("network", network),
					gd.gauntlet.Flag("state", StoreAccount), // why is this needed in gauntlet?
					gd.gauntlet.Flag("input", string(jsonInput))},
				solanaCommandError,
				RETRY_COUNT,
			)
			Expect(err).ShouldNot(HaveOccurred(), "Creating a feed for the store failed")
			OCRTransmissions := report.Data["transmissions"]

			log.Info().Msg("Set Validator Config in Store")
			report, _, err = gd.gauntlet.ExecuteAndRead(
				[]string{"store:set_validator_config",
					gd.gauntlet.Flag("network", network),
					gd.gauntlet.Flag("state", StoreAccount),
					gd.gauntlet.Flag("feed", OCRTransmissions),
					gd.gauntlet.Flag("threshold", "8000")},
				solanaCommandError,
				RETRY_COUNT,
			)
			Expect(err).ShouldNot(HaveOccurred(), "Setting the stores validator config failed")

			log.Info().Msg("Init OCR 2 Feed")
			input = map[string]interface{}{
				"minAnswer":     "0",
				"maxAnswer":     "10000000000",
				"transmissions": OCRTransmissions,
			}

			jsonInput, err = json.Marshal(input)
			Expect(err).ShouldNot(HaveOccurred(), "Marshalling the ocr2 input failed")

			// TODO: command doesn't throw an error in go if it fails
			// time.Sleep(30 * time.Second) // give time for everything else to complete
			report, _, err = gd.gauntlet.ExecuteAndRead(
				[]string{"ocr2:initialize",
					gd.gauntlet.Flag("network", network),
					gd.gauntlet.Flag("requesterAccessController", RequesterAccessController),
					gd.gauntlet.Flag("billingAccessController", BillingAccessController),
					gd.gauntlet.Flag("link", linkAddress),
					gd.gauntlet.Flag("input", string(jsonInput))},
				solanaCommandError,
				RETRY_COUNT,
			)
			Expect(err).ShouldNot(HaveOccurred(), "Initializing ocr2 failed")

			// OCRFeed := report.Data["state"]
			// StoreAuthority := report.Data["storeAuthority"]
		})
		// It("deviation_flagging_validator", func() {
		// 	Expect("abc").To(Equal("abc"))
		// 	// deviation_flagging_validator:initialize
		// })
	})

	AfterEach(func() {
		By("Tearing down the environment", func() {
			err = actions.TeardownSuite(state.Env, nil, "logs")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
