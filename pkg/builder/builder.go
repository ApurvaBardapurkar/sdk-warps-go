package builder

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/cache"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/utils"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/validator"
)

// WarpBuilder provides functionality for building and creating warps
type WarpBuilder struct {
	config     types.WarpConfig
	cache      *cache.WarpCache
	pendingWarp types.Warp
}

// NewWarpBuilder creates a new WarpBuilder instance
func NewWarpBuilder(config types.WarpConfig) *WarpBuilder {
	return &WarpBuilder{
		config: config,
		cache:  cache.NewWarpCache(),
		pendingWarp: types.Warp{
			Protocol:    utils.GetLatestProtocolIdentifier(types.WarpProtocol),
			Name:        "",
			Title:       "",
			Description: nil,
			Preview:     nil,
			Actions:     []types.WarpAction{},
		},
	}
}

// CreateInscriptionTransaction creates a transaction to inscribe a warp on the blockchain
func (b *WarpBuilder) CreateInscriptionTransaction(warp *types.Warp) (string, error) {
	if b.config.UserAddress == "" {
		return "", errors.New("WarpBuilder: user address not set")
	}

	// For now, we just return a serialized warp as a transaction payload
	// In a real implementation, this would create an actual blockchain transaction
	serialized, err := json.Marshal(warp)
	if err != nil {
		return "", err
	}

	return string(serialized), nil
}

// CreateFromRaw creates a warp from a raw JSON string
func (b *WarpBuilder) CreateFromRaw(encoded string, validate bool) (*types.Warp, error) {
	var warp types.Warp
	if err := json.Unmarshal([]byte(encoded), &warp); err != nil {
		return nil, err
	}

	if validate {
		warpValidator := validator.NewWarpValidator(b.config)
		if err := warpValidator.Validate(&warp); err != nil {
			return nil, err
		}
	}

	return utils.PrepareVars(&warp, b.config), nil
}

// CreateFromTransaction creates a warp from a transaction
func (b *WarpBuilder) CreateFromTransaction(txData string, sender string, timestamp int64, txHash string, validate bool) (*types.Warp, error) {
	warp, err := b.CreateFromRaw(txData, validate)
	if err != nil {
		return nil, err
	}

	// Add metadata
	warp.Meta = &types.WarpMeta{
		Hash:      txHash,
		Creator:   sender,
		CreatedAt: utils.FormatTimeISO8601(time.Unix(timestamp, 0)),
	}

	return warp, nil
}

// CreateFromTransactionHash creates a warp from a transaction hash
func (b *WarpBuilder) CreateFromTransactionHash(hash string, cacheConfig *types.WarpCacheConfig) (*types.Warp, error) {
	// Check cache
	if cacheConfig != nil {
		cachedWarp := b.cache.Get(cache.CacheKey.Warp(hash))
		if cachedWarp != nil {
			return cachedWarp.(*types.Warp), nil
		}
	}

	// In a real implementation, this would call the chain API to get the transaction
	// For now, we simulate it with a placeholder
	// For real implementation, use the ChainAPIURL from the config
	
	// Simplified for the example - in a real implementation, you would call the MultiversX API
	chainAPIURL := b.config.ChainAPIURL
	if chainAPIURL == "" {
		chainAPIURL = "https://api.multiversx.com"
	}

	resp, err := http.Get(chainAPIURL + "/transactions/" + hash)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get transaction")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var txResponse struct {
		Data     string `json:"data"`
		Sender   string `json:"sender"`
		Timestamp int64  `json:"timestamp"`
	}

	if err := json.Unmarshal(body, &txResponse); err != nil {
		return nil, err
	}

	warp, err := b.CreateFromTransaction(txResponse.Data, txResponse.Sender, txResponse.Timestamp, hash, false)
	if err != nil {
		return nil, err
	}

	// Cache the warp if caching is enabled
	if cacheConfig != nil && cacheConfig.TTL > 0 {
		b.cache.Set(cache.CacheKey.Warp(hash), warp, cacheConfig.TTL)
	}

	return warp, nil
}

// SetName sets the name of the pending warp
func (b *WarpBuilder) SetName(name string) *WarpBuilder {
	b.pendingWarp.Name = name
	return b
}

// SetTitle sets the title of the pending warp
func (b *WarpBuilder) SetTitle(title string) *WarpBuilder {
	b.pendingWarp.Title = title
	return b
}

// SetDescription sets the description of the pending warp
func (b *WarpBuilder) SetDescription(description string) *WarpBuilder {
	b.pendingWarp.Description = &description
	return b
}

// SetPreview sets the preview of the pending warp
func (b *WarpBuilder) SetPreview(preview string) *WarpBuilder {
	b.pendingWarp.Preview = &preview
	return b
}

// SetActions sets the actions of the pending warp
func (b *WarpBuilder) SetActions(actions []types.WarpAction) *WarpBuilder {
	b.pendingWarp.Actions = actions
	return b
}

// AddAction adds an action to the pending warp
func (b *WarpBuilder) AddAction(action types.WarpAction) *WarpBuilder {
	b.pendingWarp.Actions = append(b.pendingWarp.Actions, action)
	return b
}

// Build builds the pending warp
func (b *WarpBuilder) Build() (*types.Warp, error) {
	// Validate required fields
	if b.pendingWarp.Protocol == "" {
		return nil, errors.New("WarpBuilder: protocol is required")
	}
	if b.pendingWarp.Name == "" {
		return nil, errors.New("WarpBuilder: name is required")
	}
	if b.pendingWarp.Title == "" {
		return nil, errors.New("WarpBuilder: title is required")
	}
	if len(b.pendingWarp.Actions) == 0 {
		return nil, errors.New("WarpBuilder: actions are required")
	}

	// Validate the warp
	warpValidator := validator.NewWarpValidator(b.config)
	if err := warpValidator.Validate(&b.pendingWarp); err != nil {
		return nil, err
	}

	return &b.pendingWarp, nil
}

// GetDescriptionPreview returns a preview of the description
func (b *WarpBuilder) GetDescriptionPreview(description string, maxChars int) string {
	return utils.ToPreviewText(description, maxChars)
} 