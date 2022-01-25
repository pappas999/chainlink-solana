package solclient

import (
	"github.com/smartcontractkit/helmenv/environment"
)

// NewChainlinkSolOCRv2 returns a cluster config with Solana test validator
func NewChainlinkSolOCRv2(chainlinkNodeCount int) *environment.Config {
	return &environment.Config{
		NamespacePrefix: "chainlink-sol",
		Charts: environment.Charts{
			"solana-validator": {
				Index: 1,
				Values: map[string]interface{}{
					"sol": map[string]interface{}{
						"image": map[string]interface{}{
							"image":   "tateexon/solana-validator",
							"version": "1.9.5-1",
						},
					},
					"resources": map[string]interface{}{
						"requests": map[string]interface{}{
							"cpu":    "2000m",
							"memory": "2000Mi",
						},
						"limits": map[string]interface{}{
							"cpu":    "2500m",
							"memory": "2000Mi",
						},
					},
				},
			},
			"mockserver-config": {
				Index: 2,
			},
			"mockserver": {
				Index: 3,
			},
			"chainlink": {
				Index: 4,
				Values: map[string]interface{}{
					"replicas": chainlinkNodeCount,
					"chainlink": map[string]interface{}{
						"image": map[string]interface{}{
							"image":   "public.ecr.aws/chainlink/chainlink",
							"version": "develop.f20690e8ede0cfead9df7f808f56a14f26469aaa",
						},
					},
					"env": map[string]interface{}{
						"eth_url":                     "ws://sol:8900",
						"eth_disabled":                "true",
						"USE_LEGACY_ETH_ENV_VARS":     "false",
						"FEATURE_OFFCHAIN_REPORTING2": "true",
						"feature_external_initiators": "true",
						"P2P_NETWORKING_STACK":        "V2",
						"P2PV2_LISTEN_ADDRESSES":      "0.0.0.0:6690",
						"P2PV2_DELTA_DIAL":            "5s",
						"P2PV2_DELTA_RECONCILE":       "5s",
						"p2p_listen_port":             "0",
					},
				},
			},
		},
	}
}

// NewChainlinkSolOCRv2 returns a cluster config with Solana test validator
func NewSolanaValidator() *environment.Config {
	return &environment.Config{
		NamespacePrefix: "chainlink-sol",
		Charts: environment.Charts{
			"solana-validator": {
				Index: 1,
			},
		},
	}
}
