# Warps SDK for Go

A Go implementation of the Warps SDK, providing a seamless interface for creating and interacting with Warps on the MultiversX blockchain.

## Features

- Create and build warps with various action types
- Generate warp links and QR codes
- Detect warps from URLs and HTML content
- Search for warps in the registry
- Register aliases and brands for warps
- Validate warps against the schema
- Unified SDK interface for easy integration

## Installation

```bash
go get github.com/ApurvaBardapurkar/sdk-warps-go
```

## Quick Start

```go
package main

import (
    "fmt"
    
    "github.com/ApurvaBardapurkar/sdk-warps-go/pkg/warp"
)

func main() {
    // Create a new SDK instance with mainnet configuration
    sdk := warp.NewSDK(warp.MainnetConfig())
    
    // Set the user address
    sdk.Config.UserAddress = "erd1..." // Replace with your wallet address
    
    // Create a warp
    sdk.Builder.SetName("my-warp")
    sdk.Builder.SetTitle("My First Warp")
    sdk.Builder.SetDescription("This is my first warp created with the Go SDK")
    
    // Build the warp
    warp, err := sdk.Builder.Build()
    if err != nil {
        fmt.Println("Error building warp:", err)
        return
    }
    
    // Generate a warp link
    url := sdk.Link.Build(types.HashIDType, "your-hash-id")
    fmt.Println("Warp link:", url)
}
```

## Usage

### Unified SDK

The unified SDK provides a single entry point to all SDK components:

```go
// Create a new SDK instance with mainnet configuration
sdk := warp.NewSDK(warp.MainnetConfig())

// Access the components
sdk.Link       // WarpLink for generating and detecting warp links
sdk.Builder    // WarpBuilder for creating and building warps
sdk.Registry   // WarpRegistry for interacting with the warp registry
sdk.Validator  // WarpValidator for validating warps
```

### Configuration

You can create a configuration for different environments:

```go
// Mainnet configuration
config := warp.MainnetConfig()

// Testnet configuration
config := warp.TestnetConfig()

// Devnet configuration
config := warp.DevnetConfig()

// Custom configuration
config := types.WarpConfig{
    Env:         types.Mainnet,
    ClientURL:   "https://usewarp.to/to",
    UserAddress: "erd1...", // Your wallet address
    ChainAPIURL: "https://api.multiversx.com",
}
```

### Creating a Warp

```go
// Create a warp
sdk.Builder.SetName("my-warp")
sdk.Builder.SetTitle("My First Warp")
sdk.Builder.SetDescription("This is my first warp created with the Go SDK")

// Add a contract action
contractAction := types.WarpContractAction{
    Type:        types.ContractActionType,
    Label:       "Deploy Contract",
    Description: stringPtr("Deploys a new smart contract"),
    Address:     "erd1...", // Contract address
    Func:        stringPtr("deployContract"),
    Args:        []string{"0x01", "0x02"},
    GasLimit:    10000000,
}
sdk.Builder.AddAction(contractAction)

// Build the warp
warp, err := sdk.Builder.Build()
if err != nil {
    fmt.Println("Error building warp:", err)
    return
}

// Create a transaction for the warp
tx, err := sdk.Builder.CreateInscriptionTransaction(warp)
if err != nil {
    fmt.Println("Error creating transaction:", err)
    return
}
```

### Generating Warp Links

```go
// Generate a warp link
url := sdk.Link.Build(types.HashIDType, "your-hash-id")
fmt.Println("Warp link:", url)

// Generate a QR code
qrCode, err := sdk.Link.GenerateQRCode(types.HashIDType, "your-hash-id", 256)
if err != nil {
    fmt.Println("Error generating QR code:", err)
    return
}

// Save the QR code to a file
err = os.WriteFile("warp-qr.png", qrCode, 0644)
if err != nil {
    fmt.Println("Error saving QR code:", err)
    return
}
```

### Detecting Warps

```go
// Detect a warp from a URL
url := "https://usewarp.to/to?warp=hash:your-hash-id"
result, err := sdk.Link.Detect(url)
if err != nil {
    fmt.Println("Error detecting warp:", err)
    return
}

if result.Match {
    fmt.Println("Detected warp:", result.Warp)
} else {
    fmt.Println("No warp detected")
}

// Detect warps from HTML content
html := "<a href='https://usewarp.to/to?warp=hash:your-hash-id'>Click here</a>"
htmlResult, err := sdk.Link.DetectFromHTML(html)
if err != nil {
    fmt.Println("Error detecting warps from HTML:", err)
    return
}

if htmlResult.Match {
    fmt.Println("Detected warps:", len(htmlResult.Results))
    for _, warpResult := range htmlResult.Results {
        fmt.Println("  URL:", warpResult.URL)
        fmt.Println("  Warp:", warpResult.Warp)
    }
}
```

### Searching for Warps

```go
// Search for warps
results, err := sdk.Registry.Search("example")
if err != nil {
    fmt.Println("Error searching for warps:", err)
    return
}

fmt.Println("Found warps:", len(results.Hits))
for _, hit := range results.Hits {
    fmt.Println("  Title:", hit.Title)
    fmt.Println("  Description:", hit.Description)
}
```

### Registering Aliases and Brands

```go
// Register an alias for a warp
txHash, err := sdk.Registry.RegisterAlias("your-hash-id", "my-alias")
if err != nil {
    fmt.Println("Error registering alias:", err)
    return
}
fmt.Println("Alias registered, transaction hash:", txHash)

// Register a brand
brand := &types.Brand{
    Protocol:    "brand-0.0.2",
    Name:        "My Brand",
    Description: "My brand description",
    Logo:        "https://example.com/logo.png",
}
txHash, err = sdk.Registry.RegisterBrand(brand)
if err != nil {
    fmt.Println("Error registering brand:", err)
    return
}
fmt.Println("Brand registered, transaction hash:", txHash)
```

## Examples

The SDK includes several examples to help you get started:

- **Basic Example**: Demonstrates basic usage of individual components
- **Unified Example**: Demonstrates comprehensive usage of the unified SDK

To run the examples:

```bash
# Run the basic example
cd examples/basic
go run main.go

# Run the unified example
cd examples/unified
go run main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.