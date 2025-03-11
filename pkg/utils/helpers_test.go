package utils

import (
	"testing"
	"time"

	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
)

func TestGetChainID(t *testing.T) {
	tests := []struct {
		name     string
		env      types.ChainEnv
		expected string
	}{
		{"Mainnet", types.Mainnet, "1"},
		{"Testnet", types.Testnet, "T"},
		{"Devnet", types.Devnet, "D"},
		{"Unknown", "unknown", "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetChainID(tt.env)
			if result != tt.expected {
				t.Errorf("GetChainID(%s) = %s, expected %s", tt.env, result, tt.expected)
			}
		})
	}
}

func TestGetLatestProtocolIdentifier(t *testing.T) {
	tests := []struct {
		name     string
		protocol types.ProtocolName
		expected string
	}{
		{"Warp", types.WarpProtocol, "warp-0.0.2"},
		{"Brand", types.BrandProtocol, "brand-0.0.2"},
		{"ABI", types.AbiProtocol, "abi-0.0.2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetLatestProtocolIdentifier(tt.protocol)
			if result != tt.expected {
				t.Errorf("GetLatestProtocolIdentifier(%s) = %s, expected %s", tt.protocol, result, tt.expected)
			}
		})
	}
}

func TestToPreviewText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		maxChars int
		expected string
	}{
		{"Empty", "", 10, ""},
		{"Short", "Hello", 10, "Hello"},
		{"Exact", "Hello World", 11, "Hello World"},
		{"Long", "Hello World, this is a long text", 15, "Hello World..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPreviewText(tt.text, tt.maxChars)
			if result != tt.expected {
				t.Errorf("ToPreviewText(%s, %d) = %s, expected %s", tt.text, tt.maxChars, result, tt.expected)
			}
		})
	}
}

func TestFormatTimeISO8601(t *testing.T) {
	// Create a fixed time for testing
	fixedTime := time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC)
	expected := "2023-01-02T03:04:05Z"

	result := FormatTimeISO8601(fixedTime)
	if result != expected {
		t.Errorf("FormatTimeISO8601() = %s, expected %s", result, expected)
	}
}

func TestParseTimeISO8601(t *testing.T) {
	// Create a fixed time string for testing
	timeStr := "2023-01-02T03:04:05Z"
	expected := time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC)

	result, err := ParseTimeISO8601(timeStr)
	if err != nil {
		t.Errorf("ParseTimeISO8601() error = %v", err)
		return
	}

	if !result.Equal(expected) {
		t.Errorf("ParseTimeISO8601() = %v, expected %v", result, expected)
	}
}

func TestGetInfoFromPrefixedIdentifier(t *testing.T) {
	tests := []struct {
		name       string
		identifier string
		wantType   types.WarpIDType
		wantID     string
		wantNil    bool
	}{
		{"Hash with prefix", "hash:123456789abcdef", types.HashIDType, "123456789abcdef", false},
		{"Alias with prefix", "alias:my-alias", types.AliasIDType, "my-alias", false},
		{"Hash without prefix", "123456789abcdef123456789abcdef123456789abcdef123456789abcdef", types.HashIDType, "123456789abcdef123456789abcdef123456789abcdef123456789abcdef", false},
		{"Alias without prefix", "my-alias", types.AliasIDType, "my-alias", false},
		{"Invalid", "invalid:format:", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetInfoFromPrefixedIdentifier(tt.identifier)
			
			if tt.wantNil {
				if result != nil {
					t.Errorf("GetInfoFromPrefixedIdentifier(%s) = %v, expected nil", tt.identifier, result)
				}
				return
			}
			
			if result == nil {
				t.Errorf("GetInfoFromPrefixedIdentifier(%s) = nil, expected non-nil", tt.identifier)
				return
			}
			
			if result.Type != tt.wantType {
				t.Errorf("GetInfoFromPrefixedIdentifier(%s).Type = %s, expected %s", tt.identifier, result.Type, tt.wantType)
			}
			
			if result.ID != tt.wantID {
				t.Errorf("GetInfoFromPrefixedIdentifier(%s).ID = %s, expected %s", tt.identifier, result.ID, tt.wantID)
			}
		})
	}
} 