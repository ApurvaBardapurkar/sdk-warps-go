package constants

// WarpConstants contains all the constants used in the SDK
var WarpConstants = struct {
	HTTPProtocolPrefix     string
	IdentifierParamName    string
	IdentifierParamSeparator string
	IdentifierType         struct {
		Alias string
		Hash  string
	}
	ArgParamsSeparator     string
	ArgCompositeSeparator  string
	EGLD                   struct {
		Identifier string
		DisplayName string
		Decimals    int
	}
}{
	HTTPProtocolPrefix:     "http",
	IdentifierParamName:    "warp",
	IdentifierParamSeparator: ":",
	IdentifierType: struct {
		Alias string
		Hash  string
	}{
		Alias: "alias",
		Hash:  "hash",
	},
	ArgParamsSeparator:    ":",
	ArgCompositeSeparator: "|",
	EGLD: struct {
		Identifier string
		DisplayName string
		Decimals    int
	}{
		Identifier: "EGLD",
		DisplayName: "eGold",
		Decimals:    18,
	},
} 