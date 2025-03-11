package main

import (
	"fmt"
	"os"

	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/warp"
)

func main() {
	fmt.Println("Warps SDK for Go")
	fmt.Println("===============")
	fmt.Println("\nThis is a library package. To use it in your project, import it as:")
	fmt.Println("  import \"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/warp\"")
	fmt.Println("\nFor examples, see the examples directory:")
	fmt.Println("  - examples/basic: Basic usage of individual components")
	fmt.Println("  - examples/unified: Comprehensive example using the unified SDK")
	
	// Create a new SDK instance with mainnet configuration
	sdk := warp.NewSDK(warp.MainnetConfig())
	
	// Print the SDK version and configuration
	fmt.Println("\nSDK Configuration:")
	fmt.Printf("  Environment: %s\n", sdk.Config.Env)
	fmt.Printf("  Client URL: %s\n", sdk.Config.ClientURL)
	fmt.Printf("  Chain API URL: %s\n", sdk.Config.ChainAPIURL)
	
	fmt.Println("\nFor more information, visit: https://github.com/ApurvaBardapurkar/sdk-warps-go")
	
	os.Exit(0)
} 