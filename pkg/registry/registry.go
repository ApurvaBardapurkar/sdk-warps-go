package registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/cache"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/core"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
)

// RegistryResult represents the result of a registry query
type RegistryResult struct {
	RegistryInfo *types.RegistryInfo `json:"registryInfo"`
	Brand        *types.Brand        `json:"brand"`
}

// WarpRegistry provides functionality for interacting with the warp registry
type WarpRegistry struct {
	config types.WarpConfig
	cache  *cache.WarpCache
}

// NewWarpRegistry creates a new WarpRegistry instance
func NewWarpRegistry(config types.WarpConfig) *WarpRegistry {
	return &WarpRegistry{
		config: config,
		cache:  cache.NewWarpCache(),
	}
}

// GetInfoByHash gets registry information by transaction hash
func (r *WarpRegistry) GetInfoByHash(hash string) (*RegistryResult, error) {
	// Check cache
	cachedRegistryInfo := r.cache.Get(cache.CacheKey.RegistryInfo(hash))
	if cachedRegistryInfo != nil {
		return cachedRegistryInfo.(*RegistryResult), nil
	}

	// In a real implementation, this would call the registry contract
	// For this example, we'll just make a simulated HTTP request
	contractAddress := r.config.RegistryContract
	if contractAddress == "" {
		contractAddress = core.Config.DefaultRegistryContract(r.config.Env)
	}

	// Example API call to get registry info
	chainAPIURL := r.config.ChainAPIURL
	if chainAPIURL == "" {
		chainAPIURL = core.Config.DefaultChainAPIURL(r.config.Env)
	}

	// Construct a URL for querying registry info
	// This is a simplified example; in reality, you would call the contract directly
	apiURL := fmt.Sprintf("%s/vm-values/query", chainAPIURL)

	// Prepare the request body - this is a simplified example
	requestBody, err := json.Marshal(map[string]interface{}{
		"scAddress": contractAddress,
		"funcName":  "getWarpByHash",
		"args":      []string{hash},
	})
	if err != nil {
		return nil, err
	}

	// Make the HTTP request
	resp, err := http.Post(apiURL, "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get registry info: %s", resp.Status)
	}

	// In a real implementation, you would parse the response
	// For this example, we'll create a simulated response
	registryInfo := &types.RegistryInfo{
		Hash:      hash,
		Alias:     nil,
		Trust:     types.Unverified,
		Creator:   "erd1...",
		CreatedAt: 1625097600,
		Brand:     nil,
		Upgrade:   nil,
	}

	// Create the result
	result := &RegistryResult{
		RegistryInfo: registryInfo,
		Brand:        nil,
	}

	// Cache the result
	if r.config.CacheTTL > 0 {
		r.cache.Set(cache.CacheKey.RegistryInfo(hash), result, r.config.CacheTTL)
	}

	return result, nil
}

// GetInfoByAlias gets registry information by alias
func (r *WarpRegistry) GetInfoByAlias(alias string) (*RegistryResult, error) {
	// Check cache
	cachedRegistryInfo := r.cache.Get(cache.CacheKey.RegistryInfo(alias))
	if cachedRegistryInfo != nil {
		return cachedRegistryInfo.(*RegistryResult), nil
	}

	// In a real implementation, this would call the registry contract
	// For this example, we'll just make a simulated HTTP request
	contractAddress := r.config.RegistryContract
	if contractAddress == "" {
		contractAddress = core.Config.DefaultRegistryContract(r.config.Env)
	}

	// Example API call to get registry info
	chainAPIURL := r.config.ChainAPIURL
	if chainAPIURL == "" {
		chainAPIURL = core.Config.DefaultChainAPIURL(r.config.Env)
	}

	// Construct a URL for querying registry info
	// This is a simplified example; in reality, you would call the contract directly
	apiURL := fmt.Sprintf("%s/vm-values/query", chainAPIURL)

	// Prepare the request body - this is a simplified example
	requestBody, err := json.Marshal(map[string]interface{}{
		"scAddress": contractAddress,
		"funcName":  "getWarpByAlias",
		"args":      []string{alias},
	})
	if err != nil {
		return nil, err
	}

	// Make the HTTP request
	resp, err := http.Post(apiURL, "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get registry info: %s", resp.Status)
	}

	// In a real implementation, you would parse the response
	// For this example, we'll create a simulated response
	hash := "123456789abcdef"
	aliasPtr := alias
	registryInfo := &types.RegistryInfo{
		Hash:      hash,
		Alias:     &aliasPtr,
		Trust:     types.Unverified,
		Creator:   "erd1...",
		CreatedAt: 1625097600,
		Brand:     nil,
		Upgrade:   nil,
	}

	// Create the result
	result := &RegistryResult{
		RegistryInfo: registryInfo,
		Brand:        nil,
	}

	// Cache the result
	if r.config.CacheTTL > 0 {
		r.cache.Set(cache.CacheKey.RegistryInfo(alias), result, r.config.CacheTTL)
	}

	return result, nil
}

// Search searches the registry for warps
func (r *WarpRegistry) Search(query string) (*types.WarpSearchResult, error) {
	// In a real implementation, this would call the index API
	indexURL := r.config.IndexURL
	if indexURL == "" {
		indexURL = core.Config.DefaultIndexURL(r.config.Env)
	}

	indexSearchParamName := r.config.IndexSearchParamName
	if indexSearchParamName == "" {
		indexSearchParamName = core.Config.DefaultIndexSearchParamName
	}

	// Construct the search URL
	searchURL := fmt.Sprintf("%s/search?%s=%s", 
		indexURL, 
		indexSearchParamName, 
		url.QueryEscape(query))

	// Add API key if available
	if r.config.IndexAPIKey != "" {
		searchURL = fmt.Sprintf("%s&apiKey=%s", searchURL, r.config.IndexAPIKey)
	}

	// Make the HTTP request
	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to search registry: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result types.WarpSearchResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// RegisterAlias registers an alias for a warp
func (r *WarpRegistry) RegisterAlias(hash string, alias string) (string, error) {
	// In a real implementation, this would call the registry contract
	if r.config.UserAddress == "" {
		return "", errors.New("WarpRegistry: user address not set")
	}

	contractAddress := r.config.RegistryContract
	if contractAddress == "" {
		contractAddress = core.Config.DefaultRegistryContract(r.config.Env)
	}

	// For this example, we'll just return a simulated transaction hash
	return "tx-hash-for-register-alias", nil
}

// RegisterBrand registers a brand for a warp
func (r *WarpRegistry) RegisterBrand(brand *types.Brand) (string, error) {
	// In a real implementation, this would call the registry contract
	if r.config.UserAddress == "" {
		return "", errors.New("WarpRegistry: user address not set")
	}

	contractAddress := r.config.RegistryContract
	if contractAddress == "" {
		contractAddress = core.Config.DefaultRegistryContract(r.config.Env)
	}

	// Validate the brand
	if brand.Protocol == "" {
		return "", errors.New("WarpRegistry: brand protocol is required")
	}
	if brand.Name == "" {
		return "", errors.New("WarpRegistry: brand name is required")
	}
	if brand.Description == "" {
		return "", errors.New("WarpRegistry: brand description is required")
	}
	if brand.Logo == "" {
		return "", errors.New("WarpRegistry: brand logo is required")
	}

	// For this example, we'll just return a simulated transaction hash
	return "tx-hash-for-register-brand", nil
} 