package smoke

import (
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	relayUtils "github.com/smartcontractkit/chainlink-relay/ops/utils"
	"github.com/smartcontractkit/chainlink-solana/tests/e2e/solclient"
	"github.com/smartcontractkit/helmenv/environment"
	"github.com/smartcontractkit/helmenv/tools"
	"github.com/smartcontractkit/integrations-framework/actions"
	"github.com/smartcontractkit/integrations-framework/client"
)

type Deployer struct {
	gauntlet relayUtils.Gauntlet
	network  string
	Account  map[int]string
}

var _ = Describe("Gauntlet Testing @gauntlet", func() {
	var (
		e              *environment.Environment
		chainlinkNodes []client.Chainlink
		// cd             contracts.ContractDeployer
		gd       Deployer
		gauntlet relayUtils.Gauntlet
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
			mockserver, err = client.NewMockServerClientFromEnv(e)
			Expect(err).ShouldNot(HaveOccurred())
			chainlinkNodes, err = client.NewChainlinkClients(e)
			Expect(err).ShouldNot(HaveOccurred())
			_, nkb, err = DefaultOffChainConfigParamsFromNodes(chainlinkNodes)
			// ocConfig, nkb, err = DefaultOffChainConfigParamsFromNodes(chainlinkNodes)
			Expect(err).ShouldNot(HaveOccurred())
		})
		By("Setup Gauntlet", func() {
			_, err := exec.LookPath("yarn")
			Expect(err).ShouldNot(HaveOccurred())
			exec.Command("").
				os.Setenv("SKIP_PROMPTS", "true")

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
			gauntlet, err = relayUtils.NewGauntlet(gauntletBin)
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
			// lt, err := cd.DeployLinkTokenContract()
			// Expect(err).ShouldNot(HaveOccurred())
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
			network := "blarg"
			networkConfig := map[string]string{
				"NETWORK":     "local",
				"NODE_URL":    fmt.Sprintf("'%s'", chainlinkNodes[0].URL()),
				"PRIVATE_KEY": "[4,198,110,68,205,35,72,68,186,229,85,4,205,206,147,113,153,7,206,234,55,216,227,86,133,186,206,65,6,175,218,9,80,251,248,55,192,181,151,105,15,181,169,175,196,47,24,179,25,173,203,190,160,153,152,66,24,155,233,243,140,153,66,143]",
				// "CHAIN_ID":    "",
			}
			err = WriteNetworkConfigMap(fmt.Sprintf("networks/.env.%s", network), networkConfig)
			Expect(err).ShouldNot(HaveOccurred())

			log.Debug().Msg("Deploying LINK Token...")
			err = gd.gauntlet.ExecCommand(
				"token:deploy",
				gd.gauntlet.Flag("network", network),
			)
			Expect(err).ShouldNot(HaveOccurred())

			_, err = gd.gauntlet.ReadCommandReport()
			Expect(err).ShouldNot(HaveOccurred())
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
