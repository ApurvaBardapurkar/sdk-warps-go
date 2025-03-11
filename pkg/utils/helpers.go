package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/constants"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
)

// GetChainID returns the chain ID for the specified environment
func GetChainID(env types.ChainEnv) string {
	switch env {
	case types.Mainnet:
		return "1"
	case types.Testnet:
		return "T"
	case types.Devnet:
		return "D"
	default:
		return "1"
	}
}

// GetLatestProtocolIdentifier returns the latest protocol identifier for the specified protocol
func GetLatestProtocolIdentifier(protocol types.ProtocolName) string {
	return fmt.Sprintf("%s-0.0.2", protocol)
}

// ToPreviewText converts a string to a preview text limited to the specified number of characters
func ToPreviewText(text string, maxChars int) string {
	if text == "" {
		return ""
	}

	if len(text) <= maxChars {
		return text
	}

	return fmt.Sprintf("%s...", text[:maxChars-3])
}

// GetInfoFromPrefixedIdentifier extracts the identifier type and ID from a prefixed identifier
func GetInfoFromPrefixedIdentifier(identifier string) *struct {
	Type types.WarpIDType
	ID   string
} {
	// Check if it has a known prefix
	if strings.HasPrefix(identifier, constants.WarpConstants.IdentifierType.Hash+constants.WarpConstants.IdentifierParamSeparator) {
		return &struct {
			Type types.WarpIDType
			ID   string
		}{
			Type: types.HashIDType,
			ID:   strings.TrimPrefix(identifier, constants.WarpConstants.IdentifierType.Hash+constants.WarpConstants.IdentifierParamSeparator),
		}
	}

	if strings.HasPrefix(identifier, constants.WarpConstants.IdentifierType.Alias+constants.WarpConstants.IdentifierParamSeparator) {
		return &struct {
			Type types.WarpIDType
			ID   string
		}{
			Type: types.AliasIDType,
			ID:   strings.TrimPrefix(identifier, constants.WarpConstants.IdentifierType.Alias+constants.WarpConstants.IdentifierParamSeparator),
		}
	}

	// If no prefix, assume it's a transaction hash
	hashPattern := regexp.MustCompile(`^[a-f0-9]{64}$`)
	if hashPattern.MatchString(identifier) {
		return &struct {
			Type types.WarpIDType
			ID   string
		}{
			Type: types.HashIDType,
			ID:   identifier,
		}
	}

	// If it looks like an alias (simple string), treat it as an alias
	aliasPattern := regexp.MustCompile(`^[a-zA-Z0-9-_]{3,}$`)
	if aliasPattern.MatchString(identifier) {
		return &struct {
			Type types.WarpIDType
			ID   string
		}{
			Type: types.AliasIDType,
			ID:   identifier,
		}
	}

	return nil
}

// FormatTimeISO8601 formats a time.Time value as an ISO8601 string
func FormatTimeISO8601(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseTimeISO8601 parses an ISO8601 string into a time.Time value
func ParseTimeISO8601(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// PrepareVars replaces variable placeholders in a warp with values from the config
func PrepareVars(warp *types.Warp, config types.WarpConfig) *types.Warp {
	if warp == nil || warp.Vars == nil || config.Vars == nil {
		return warp
	}

	// Deep copy the warp to avoid modifying the original
	result := *warp

	// Apply the variables
	for key, value := range warp.Vars {
		if configValue, exists := config.Vars[string(key)]; exists {
			warp.Vars[key] = configValue
		}
	}

	return &result
} 