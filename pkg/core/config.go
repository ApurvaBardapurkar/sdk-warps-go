package core

import (
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
)

// Config holds the default configuration values and methods
var Config = struct {
	// SuperClientUrls contains the list of URLs that can be used as super clients
	SuperClientURLs []string

	// DefaultClientURL returns the default client URL for the specified environment
	DefaultClientURL func(env types.ChainEnv) string

	// DefaultChainAPIURL returns the default chain API URL for the specified environment
	DefaultChainAPIURL func(env types.ChainEnv) string

	// DefaultWarpSchemaURL returns the default warp schema URL for the specified environment
	DefaultWarpSchemaURL func(env types.ChainEnv) string

	// DefaultBrandSchemaURL returns the default brand schema URL for the specified environment
	DefaultBrandSchemaURL func(env types.ChainEnv) string

	// DefaultRegistryContract returns the default registry contract address for the specified environment
	DefaultRegistryContract func(env types.ChainEnv) string

	// DefaultIndexURL returns the default index URL for the specified environment
	DefaultIndexURL func(env types.ChainEnv) string

	// DefaultIndexSearchParamName returns the default index search parameter name
	DefaultIndexSearchParamName string
}{
	SuperClientURLs: []string{
		"https://warp.to",
		"https://usewarp.to",
	},
	DefaultClientURL: func(env types.ChainEnv) string {
		switch env {
		case types.Mainnet:
			return "https://usewarp.to/to"
		case types.Testnet:
			return "https://testnet.usewarp.to/to"
		case types.Devnet:
			return "https://devnet.usewarp.to/to"
		default:
			return "https://usewarp.to/to"
		}
	},
	DefaultChainAPIURL: func(env types.ChainEnv) string {
		switch env {
		case types.Mainnet:
			return "https://api.multiversx.com"
		case types.Testnet:
			return "https://testnet-api.multiversx.com"
		case types.Devnet:
			return "https://devnet-api.multiversx.com"
		default:
			return "https://api.multiversx.com"
		}
	},
	DefaultWarpSchemaURL: func(env types.ChainEnv) string {
		return "https://raw.githubusercontent.com/usewarps/schema/main/warp.schema.json"
	},
	DefaultBrandSchemaURL: func(env types.ChainEnv) string {
		return "https://raw.githubusercontent.com/usewarps/schema/main/brand.schema.json"
	},
	DefaultRegistryContract: func(env types.ChainEnv) string {
		switch env {
		case types.Mainnet:
			return "erd1qqqqqqqqqqqqqpgqt4g39q9vzm80aujs44wmvwn4kfn6czmtcrpslzjs4e"
		case types.Testnet:
			return "erd1qqqqqqqqqqqqqpgqnyj0lmcyft0v8yc5xm5vljz3y7mnkx7tcrps2jgm07"
		case types.Devnet:
			return "erd1qqqqqqqqqqqqqpgq34s7nd6sudf3jm5w44qqkpgfdxzplh4tcrps3ffxdl"
		default:
			return "erd1qqqqqqqqqqqqqpgqt4g39q9vzm80aujs44wmvwn4kfn6czmtcrpslzjs4e"
		}
	},
	DefaultIndexURL: func(env types.ChainEnv) string {
		return "https://index.usewarp.to/api"
	},
	DefaultIndexSearchParamName: "q",
} 