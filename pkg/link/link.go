package link

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/builder"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/constants"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/core"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/registry"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/utils"
	"github.com/skip2/go-qrcode"
)

// DetectionResult represents the result of a warp detection
type DetectionResult struct {
	Match        bool            `json:"match"`
	URL          string          `json:"url"`
	Warp         *types.Warp     `json:"warp"`
	RegistryInfo *types.RegistryInfo `json:"registryInfo"`
	Brand        *types.Brand    `json:"brand"`
}

// DetectionResultFromHTML represents the result of detecting warps in HTML content
type DetectionResultFromHTML struct {
	Match   bool              `json:"match"`
	Results []types.WarpHTMLResult `json:"results"`
}

// WarpLink provides functionality for generating and detecting warp links
type WarpLink struct {
	config types.WarpConfig
}

// NewWarpLink creates a new WarpLink instance
func NewWarpLink(config types.WarpConfig) *WarpLink {
	return &WarpLink{
		config: config,
	}
}

// IsValid checks if a URL is a valid warp URL
func (wl *WarpLink) IsValid(urlStr string) bool {
	if !strings.HasPrefix(urlStr, constants.WarpConstants.HTTPProtocolPrefix) {
		return false
	}

	idResult := wl.extractIdentifierInfoFromURL(urlStr)
	return idResult != nil
}

// DetectFromHTML detects warps in HTML content
func (wl *WarpLink) DetectFromHTML(content string) (*DetectionResultFromHTML, error) {
	if len(content) == 0 {
		return &DetectionResultFromHTML{
			Match:   false,
			Results: []types.WarpHTMLResult{},
		}, nil
	}

	// Find all URLs in the HTML content
	urlRegex := regexp.MustCompile(`https?://[^\s"'<>]+`)
	matches := urlRegex.FindAllString(content, -1)

	// Filter for warp links
	var warpLinks []string
	for _, link := range matches {
		if wl.IsValid(link) {
			warpLinks = append(warpLinks, link)
		}
	}

	if len(warpLinks) == 0 {
		return &DetectionResultFromHTML{
			Match:   false,
			Results: []types.WarpHTMLResult{},
		}, nil
	}

	// Detect warps for each link
	var results []types.WarpHTMLResult
	for _, link := range warpLinks {
		detection, err := wl.Detect(link)
		if err != nil {
			continue
		}

		if detection.Match && detection.Warp != nil {
			results = append(results, types.WarpHTMLResult{
				URL:  detection.URL,
				Warp: *detection.Warp,
			})
		}
	}

	return &DetectionResultFromHTML{
		Match:   len(results) > 0,
		Results: results,
	}, nil
}

// Detect detects a warp from a URL
func (wl *WarpLink) Detect(urlStr string) (*DetectionResult, error) {
	var idResult *struct {
		Type types.WarpIDType
		ID   string
	}

	if strings.HasPrefix(urlStr, constants.WarpConstants.HTTPProtocolPrefix) {
		idResult = wl.extractIdentifierInfoFromURL(urlStr)
	} else {
		idResult = utils.GetInfoFromPrefixedIdentifier(urlStr)
	}

	if idResult == nil {
		return &DetectionResult{
			Match:        false,
			URL:          urlStr,
			Warp:         nil,
			RegistryInfo: nil,
			Brand:        nil,
		}, nil
	}

	warpType := idResult.Type
	id := idResult.ID
	warpBuilder := builder.NewWarpBuilder(wl.config)
	warpRegistry := registry.NewWarpRegistry(wl.config)

	var warp *types.Warp
	var registryInfo *types.RegistryInfo
	var brand *types.Brand
	var err error

	if warpType == types.HashIDType {
		// Get the warp from the transaction hash
		warp, err = warpBuilder.CreateFromTransactionHash(id, nil)
		if err != nil {
			return &DetectionResult{
				Match:        false,
				URL:          urlStr,
				Warp:         nil,
				RegistryInfo: nil,
				Brand:        nil,
			}, err
		}

		// Try to get registry info
		registryResult, err := warpRegistry.GetInfoByHash(id)
		if err == nil && registryResult != nil {
			registryInfo = registryResult.RegistryInfo
			brand = registryResult.Brand
		}
	} else if warpType == types.AliasIDType {
		// Get the registry info by alias
		registryResult, err := warpRegistry.GetInfoByAlias(id)
		if err != nil || registryResult == nil {
			return &DetectionResult{
				Match:        false,
				URL:          urlStr,
				Warp:         nil,
				RegistryInfo: nil,
				Brand:        nil,
			}, err
		}

		registryInfo = registryResult.RegistryInfo
		brand = registryResult.Brand

		if registryInfo != nil {
			// Get the warp from the hash in registry info
			warp, err = warpBuilder.CreateFromTransactionHash(registryInfo.Hash, nil)
			if err != nil {
				return &DetectionResult{
					Match:        false,
					URL:          urlStr,
					Warp:         nil,
					RegistryInfo: nil,
					Brand:        nil,
				}, err
			}
		}
	}

	if warp == nil {
		return &DetectionResult{
			Match:        false,
			URL:          urlStr,
			Warp:         nil,
			RegistryInfo: nil,
			Brand:        nil,
		}, nil
	}

	return &DetectionResult{
		Match:        true,
		URL:          urlStr,
		Warp:         warp,
		RegistryInfo: registryInfo,
		Brand:        brand,
	}, nil
}

// Build creates a warp URL for the specified type and ID
func (wl *WarpLink) Build(idType types.WarpIDType, id string) string {
	clientURL := wl.config.ClientURL
	if clientURL == "" {
		clientURL = core.Config.DefaultClientURL(wl.config.Env)
	}

	var encodedValue string
	if idType == types.AliasIDType {
		encodedValue = url.QueryEscape(id)
	} else {
		encodedValue = url.QueryEscape(fmt.Sprintf("%s%s%s", 
			idType, 
			constants.WarpConstants.IdentifierParamSeparator, 
			id))
	}

	// Check if the client URL is a super client
	isSuperClient := false
	for _, superURL := range core.Config.SuperClientURLs {
		if strings.HasPrefix(clientURL, superURL) {
			isSuperClient = true
			break
		}
	}

	if isSuperClient {
		return fmt.Sprintf("%s/%s", clientURL, encodedValue)
	}

	return fmt.Sprintf("%s?%s=%s", clientURL, constants.WarpConstants.IdentifierParamName, encodedValue)
}

// GenerateQRCode generates a QR code for the specified warp
func (wl *WarpLink) GenerateQRCode(idType types.WarpIDType, id string, size int) ([]byte, error) {
	if size <= 0 {
		size = 256
	}

	url := wl.Build(idType, id)
	return qrcode.Encode(url, qrcode.Medium, size)
}

// extractIdentifierInfoFromURL extracts the identifier info from a URL
func (wl *WarpLink) extractIdentifierInfoFromURL(urlStr string) *struct {
	Type types.WarpIDType
	ID   string
} {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil
	}

	// Check if it's a super client
	isSuperClient := false
	for _, superURL := range core.Config.SuperClientURLs {
		if strings.HasPrefix(parsedURL.Scheme+"://"+parsedURL.Host, superURL) {
			isSuperClient = true
			break
		}
	}

	var value string
	if isSuperClient && len(parsedURL.Path) > 1 {
		// For super clients, the value might be in the path
		pathParts := strings.Split(parsedURL.Path, "/")
		if len(pathParts) > 1 {
			value = pathParts[1]
		}
	}

	// If not found in path or not a super client, look in query parameters
	if value == "" {
		value = parsedURL.Query().Get(constants.WarpConstants.IdentifierParamName)
	}

	if value == "" {
		return nil
	}

	// URL decode the value
	decodedValue, err := url.QueryUnescape(value)
	if err != nil {
		return nil
	}

	// Get the identifier info
	return utils.GetInfoFromPrefixedIdentifier(decodedValue)
} 