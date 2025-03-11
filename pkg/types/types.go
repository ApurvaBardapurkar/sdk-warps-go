package types

import (
	"fmt"
	"time"
)

// ChainEnv represents the blockchain environment
type ChainEnv string

const (
	Mainnet ChainEnv = "mainnet"
	Testnet ChainEnv = "testnet"
	Devnet  ChainEnv = "devnet"
)

// ProtocolName represents the protocol type
type ProtocolName string

const (
	WarpProtocol  ProtocolName = "warp"
	BrandProtocol ProtocolName = "brand"
	AbiProtocol   ProtocolName = "abi"
)

// WarpConfig represents the configuration for the SDK
type WarpConfig struct {
	Env                  ChainEnv          `json:"env"`
	ClientURL            string            `json:"clientUrl,omitempty"`
	CurrentURL           string            `json:"currentUrl,omitempty"`
	UserAddress          string            `json:"userAddress,omitempty"`
	ChainAPIURL          string            `json:"chainApiUrl,omitempty"`
	WarpSchemaURL        string            `json:"warpSchemaUrl,omitempty"`
	BrandSchemaURL       string            `json:"brandSchemaUrl,omitempty"`
	CacheTTL             int               `json:"cacheTtl,omitempty"`
	RegistryContract     string            `json:"registryContract,omitempty"`
	IndexURL             string            `json:"indexUrl,omitempty"`
	IndexAPIKey          string            `json:"indexApiKey,omitempty"`
	IndexSearchParamName string            `json:"indexSearchParamName,omitempty"`
	Vars                 map[string]string `json:"vars,omitempty"`
}

// WarpCacheConfig represents cache configuration
type WarpCacheConfig struct {
	TTL int `json:"ttl,omitempty"`
}

// TrustStatus represents the verification status of a warp
type TrustStatus string

const (
	Unverified  TrustStatus = "unverified"
	Verified    TrustStatus = "verified"
	Blacklisted TrustStatus = "blacklisted"
)

// RegistryInfo represents information about a warp in the registry
type RegistryInfo struct {
	Hash      string      `json:"hash"`
	Alias     *string     `json:"alias"`
	Trust     TrustStatus `json:"trust"`
	Creator   string      `json:"creator"`
	CreatedAt int64       `json:"createdAt"`
	Brand     *string     `json:"brand"`
	Upgrade   *string     `json:"upgrade"`
}

// WarpIDType represents the type of identifier for a warp
type WarpIDType string

const (
	HashIDType  WarpIDType = "hash"
	AliasIDType WarpIDType = "alias"
)

// WarpVarPlaceholder is a type for variable placeholders in warps
type WarpVarPlaceholder string

// Warp represents a complete warp definition
type Warp struct {
	Protocol    string                       `json:"protocol"`
	Name        string                       `json:"name"`
	Title       string                       `json:"title"`
	Description *string                      `json:"description"`
	Bot         *string                      `json:"bot,omitempty"`
	Preview     *string                      `json:"preview,omitempty"`
	Vars        map[WarpVarPlaceholder]string `json:"vars,omitempty"`
	Actions     []WarpAction                 `json:"actions"`
	Next        *string                      `json:"next,omitempty"`
	Meta        *WarpMeta                    `json:"meta,omitempty"`
}

// WarpMeta represents metadata about a warp
type WarpMeta struct {
	Hash      string `json:"hash"`
	Creator   string `json:"creator"`
	CreatedAt string `json:"createdAt"` // ISO8601 timestamp
}

// WarpActionType represents the type of action
type WarpActionType string

const (
	TransferActionType WarpActionType = "transfer"
	ContractActionType WarpActionType = "contract"
	QueryActionType    WarpActionType = "query"
	CollectActionType  WarpActionType = "collect"
	LinkActionType     WarpActionType = "link"
)

// WarpAction is an interface for all action types
type WarpAction interface {
	GetType() WarpActionType
	GetLabel() string
	GetDescription() *string
	GetNext() *string
}

// WarpTransferAction represents a transfer action
type WarpTransferAction struct {
	Type        WarpActionType              `json:"type"`
	Label       string                      `json:"label"`
	Description *string                     `json:"description,omitempty"`
	Address     *string                     `json:"address,omitempty"`
	Args        []string                    `json:"args,omitempty"`
	Value       *string                     `json:"value,omitempty"`
	Transfers   []WarpContractActionTransfer `json:"transfers,omitempty"`
	Inputs      []WarpActionInput           `json:"inputs,omitempty"`
	Next        *string                     `json:"next,omitempty"`
}

// GetType returns the action type
func (a WarpTransferAction) GetType() WarpActionType {
	return a.Type
}

// GetLabel returns the action label
func (a WarpTransferAction) GetLabel() string {
	return a.Label
}

// GetDescription returns the action description
func (a WarpTransferAction) GetDescription() *string {
	return a.Description
}

// GetNext returns the next action reference
func (a WarpTransferAction) GetNext() *string {
	return a.Next
}

// WarpContractAction represents a contract action
type WarpContractAction struct {
	Type        WarpActionType              `json:"type"`
	Label       string                      `json:"label"`
	Description *string                     `json:"description,omitempty"`
	Address     string                      `json:"address"`
	Func        *string                     `json:"func"`
	Args        []string                    `json:"args"`
	Value       *string                     `json:"value,omitempty"`
	GasLimit    int                         `json:"gasLimit"`
	Transfers   []WarpContractActionTransfer `json:"transfers,omitempty"`
	Inputs      []WarpActionInput           `json:"inputs,omitempty"`
	Next        *string                     `json:"next,omitempty"`
}

// GetType returns the action type
func (a WarpContractAction) GetType() WarpActionType {
	return a.Type
}

// GetLabel returns the action label
func (a WarpContractAction) GetLabel() string {
	return a.Label
}

// GetDescription returns the action description
func (a WarpContractAction) GetDescription() *string {
	return a.Description
}

// GetNext returns the next action reference
func (a WarpContractAction) GetNext() *string {
	return a.Next
}

// WarpContractActionTransfer represents a token transfer in a contract action
type WarpContractActionTransfer struct {
	Token  string  `json:"token"`
	Nonce  *int    `json:"nonce,omitempty"`
	Amount *string `json:"amount,omitempty"`
}

// WarpLinkAction represents a link action
type WarpLinkAction struct {
	Type        WarpActionType    `json:"type"`
	Label       string            `json:"label"`
	Description *string           `json:"description,omitempty"`
	URL         string            `json:"url"`
	Inputs      []WarpActionInput `json:"inputs,omitempty"`
	Next        *string           `json:"next,omitempty"`
}

// GetType returns the action type
func (a WarpLinkAction) GetType() WarpActionType {
	return a.Type
}

// GetLabel returns the action label
func (a WarpLinkAction) GetLabel() string {
	return a.Label
}

// GetDescription returns the action description
func (a WarpLinkAction) GetDescription() *string {
	return a.Description
}

// GetNext returns the next action reference
func (a WarpLinkAction) GetNext() *string {
	return a.Next
}

// WarpQueryAction represents a query action
type WarpQueryAction struct {
	Type        WarpActionType    `json:"type"`
	Label       string            `json:"label"`
	Description *string           `json:"description,omitempty"`
	Address     string            `json:"address"`
	Func        string            `json:"func"`
	Args        []string          `json:"args"`
	ABI         *string           `json:"abi,omitempty"`
	Inputs      []WarpActionInput `json:"inputs,omitempty"`
	Next        *string           `json:"next,omitempty"`
}

// GetType returns the action type
func (a WarpQueryAction) GetType() WarpActionType {
	return a.Type
}

// GetLabel returns the action label
func (a WarpQueryAction) GetLabel() string {
	return a.Label
}

// GetDescription returns the action description
func (a WarpQueryAction) GetDescription() *string {
	return a.Description
}

// GetNext returns the next action reference
func (a WarpQueryAction) GetNext() *string {
	return a.Next
}

// RequestMethod represents HTTP methods
type RequestMethod string

const (
	GET  RequestMethod = "GET"
	POST RequestMethod = "POST"
)

// WarpCollectAction represents a collect action
type WarpCollectAction struct {
	Type        WarpActionType    `json:"type"`
	Label       string            `json:"label"`
	Description *string           `json:"description,omitempty"`
	Destination struct {
		URL     string            `json:"url"`
		Method  RequestMethod     `json:"method"`
		Headers map[string]string `json:"headers"`
	} `json:"destination"`
	Inputs []WarpActionInput `json:"inputs,omitempty"`
	Next   *string           `json:"next,omitempty"`
}

// GetType returns the action type
func (a WarpCollectAction) GetType() WarpActionType {
	return a.Type
}

// GetLabel returns the action label
func (a WarpCollectAction) GetLabel() string {
	return a.Label
}

// GetDescription returns the action description
func (a WarpCollectAction) GetDescription() *string {
	return a.Description
}

// GetNext returns the next action reference
func (a WarpCollectAction) GetNext() *string {
	return a.Next
}

// WarpActionInputSource represents the source of input
type WarpActionInputSource string

const (
	FieldSource WarpActionInputSource = "field"
	QuerySource WarpActionInputSource = "query"
)

// BaseWarpActionInputType represents basic input types
type BaseWarpActionInputType string

const (
	StringInputType  BaseWarpActionInputType = "string"
	Uint8InputType   BaseWarpActionInputType = "uint8"
	Uint16InputType  BaseWarpActionInputType = "uint16"
	Uint32InputType  BaseWarpActionInputType = "uint32"
	Uint64InputType  BaseWarpActionInputType = "uint64"
	BigUintInputType BaseWarpActionInputType = "biguint"
	BoolInputType    BaseWarpActionInputType = "bool"
	AddressInputType BaseWarpActionInputType = "address"
	TokenInputType   BaseWarpActionInputType = "token"
	CodeMetaInputType BaseWarpActionInputType = "codemeta"
	HexInputType     BaseWarpActionInputType = "hex"
	EsdtInputType    BaseWarpActionInputType = "esdt"
	NftInputType     BaseWarpActionInputType = "nft"
)

// WarpActionInputType is a string representation of input types
type WarpActionInputType string

// WarpActionInputPosition represents where the input is used
type WarpActionInputPosition string

const (
	ReceiverPosition WarpActionInputPosition = "receiver"
	ValuePosition    WarpActionInputPosition = "value"
	TransferPosition WarpActionInputPosition = "transfer"
)

// ArgPosition creates a position for the specified argument index
func ArgPosition(index int) WarpActionInputPosition {
	return WarpActionInputPosition(fmt.Sprintf("arg:%d", index))
}

// WarpActionInputModifier represents input modifiers
type WarpActionInputModifier string

const (
	ScaleModifier WarpActionInputModifier = "scale"
)

// WarpActionInput represents an input for an action
type WarpActionInput struct {
	Name              string                  `json:"name"`
	As                *string                 `json:"as,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	Bot               *string                 `json:"bot,omitempty"`
	Type              WarpActionInputType     `json:"type"`
	Position          WarpActionInputPosition `json:"position"`
	Source            WarpActionInputSource   `json:"source"`
	Required          *bool                   `json:"required,omitempty"`
	Min               interface{}             `json:"min,omitempty"` // Can be number or WarpVarPlaceholder
	Max               interface{}             `json:"max,omitempty"` // Can be number or WarpVarPlaceholder
	Pattern           *string                 `json:"pattern,omitempty"`
	PatternDescription *string                `json:"patternDescription,omitempty"`
	Options           []string                `json:"options,omitempty"`
	Modifier          *string                 `json:"modifier,omitempty"`
}

// WarpActionExecutionResult represents the result of executing an action
type WarpActionExecutionResult struct {
	Action WarpAction `json:"action"`
	User   struct {
		Address string `json:"address"`
	} `json:"user"`
	Tx *string `json:"tx,omitempty"`
}

// WarpContract represents a smart contract
type WarpContract struct {
	Address  string `json:"address"`
	Owner    string `json:"owner"`
	Verified bool   `json:"verified"`
}

// WarpContractVerification represents contract verification info
type WarpContractVerification struct {
	CodeHash string      `json:"codeHash"`
	ABI      interface{} `json:"abi"`
}

// Brand represents brand information
type Brand struct {
	Protocol    string      `json:"protocol"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Logo        string      `json:"logo"`
	URLs        *BrandURLs  `json:"urls,omitempty"`
	Colors      *BrandColors `json:"colors,omitempty"`
	CTA         *BrandCTA   `json:"cta,omitempty"`
	Meta        *BrandMeta  `json:"meta,omitempty"`
}

// BrandURLs represents brand URLs
type BrandURLs struct {
	Web *string `json:"web,omitempty"`
}

// BrandColors represents brand colors
type BrandColors struct {
	Primary   *string `json:"primary,omitempty"`
	Secondary *string `json:"secondary,omitempty"`
}

// BrandCTA represents a call-to-action
type BrandCTA struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Label       string `json:"label"`
	URL         string `json:"url"`
}

// BrandMeta represents brand metadata
type BrandMeta struct {
	Hash      string `json:"hash"`
	Creator   string `json:"creator"`
	CreatedAt string `json:"createdAt"` // ISO8601 timestamp
}

// WarpSearchResult represents search results
type WarpSearchResult struct {
	Hits []WarpSearchHit `json:"hits"`
}

// WarpSearchHit represents a search result item
type WarpSearchHit struct {
	Hash        string `json:"hash"`
	Alias       string `json:"alias"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Preview     string `json:"preview"`
	Status      string `json:"status"`
	Category    string `json:"category"`
	Featured    bool   `json:"featured"`
}

// WarpAbi represents an ABI definition
type WarpAbi struct {
	Protocol string       `json:"protocol"`
	Content  AbiContents  `json:"content"`
	Meta     *WarpMeta    `json:"meta,omitempty"`
}

// AbiContents represents ABI contents
type AbiContents struct {
	Name               *string           `json:"name,omitempty"`
	Constructor        interface{}       `json:"constructor,omitempty"`
	UpgradeConstructor interface{}       `json:"upgradeConstructor,omitempty"`
	Endpoints          []interface{}     `json:"endpoints,omitempty"`
	Types              map[string]interface{} `json:"types,omitempty"`
	Events             []interface{}     `json:"events,omitempty"`
}

// DetectionResult represents the result of a warp detection
type DetectionResult struct {
	Match        bool         `json:"match"`
	URL          string       `json:"url"`
	Warp         *Warp        `json:"warp"`
	RegistryInfo *RegistryInfo `json:"registryInfo"`
	Brand        *Brand       `json:"brand"`
}

// DetectionResultFromHTML represents the result of detecting warps in HTML content
type DetectionResultFromHTML struct {
	Match   bool            `json:"match"`
	Results []WarpHTMLResult `json:"results"`
}

// WarpHTMLResult represents a warp result from HTML content
type WarpHTMLResult struct {
	URL  string `json:"url"`
	Warp Warp   `json:"warp"`
}