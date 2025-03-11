// Package warp provides the main entry point for the Warps SDK
package warp

import (
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/builder"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/core"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/link"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/registry"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/validator"
)

// SDK provides a unified interface to all SDK components
type SDK struct {
	Config     types.WarpConfig
	Link       *link.WarpLink
	Builder    *builder.WarpBuilder
	Registry   *registry.WarpRegistry
	Validator  *validator.WarpValidator
}

// NewSDK creates a new SDK instance with the specified configuration
func NewSDK(config types.WarpConfig) *SDK {
	return &SDK{
		Config:     config,
		Link:       link.NewWarpLink(config),
		Builder:    builder.NewWarpBuilder(config),
		Registry:   registry.NewWarpRegistry(config),
		Validator:  validator.NewWarpValidator(config),
	}
}

// DefaultConfig returns a default configuration for the specified environment
func DefaultConfig(env types.ChainEnv) types.WarpConfig {
	return types.WarpConfig{
		Env:                  env,
		ClientURL:            core.Config.DefaultClientURL(env),
		ChainAPIURL:          core.Config.DefaultChainAPIURL(env),
		WarpSchemaURL:        core.Config.DefaultWarpSchemaURL(env),
		BrandSchemaURL:       core.Config.DefaultBrandSchemaURL(env),
		RegistryContract:     core.Config.DefaultRegistryContract(env),
		IndexURL:             core.Config.DefaultIndexURL(env),
		IndexSearchParamName: core.Config.DefaultIndexSearchParamName,
		CacheTTL:             3600, // 1 hour
	}
}

// MainnetConfig returns a default configuration for the mainnet environment
func MainnetConfig() types.WarpConfig {
	return DefaultConfig(types.Mainnet)
}

// TestnetConfig returns a default configuration for the testnet environment
func TestnetConfig() types.WarpConfig {
	return DefaultConfig(types.Testnet)
}

// DevnetConfig returns a default configuration for the devnet environment
func DevnetConfig() types.WarpConfig {
	return DefaultConfig(types.Devnet)
} 