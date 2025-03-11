package main

import (
	"fmt"
	"os"

	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/warp"
)

func main() {
	// Create a new SDK instance with mainnet configuration
	sdk := warp.NewSDK(warp.MainnetConfig())

	// Set the user address
	sdk.Config.UserAddress = "erd1..." // Replace with your wallet address

	fmt.Println("Warps SDK for Go - Unified Example")
	fmt.Println("===================================")

	// Example 1: Create a warp
	fmt.Println("\nExample 1: Create a warp")
	sdk.Builder.SetName("unified-example")
	sdk.Builder.SetTitle("Unified SDK Example")
	sdk.Builder.SetDescription("This warp was created using the unified SDK")

	// Add a contract action
	contractAction := types.WarpContractAction{
		Type:        types.ContractActionType,
		Label:       "Execute Contract",
		Description: stringPtr("Executes a smart contract function"),
		Address:     "erd1qqqqqqqqqqqqqpgqt4g39q9vzm80aujs44wmvwn4kfn6czmtcrpslzjs4e",
		Func:        stringPtr("execute"),
		Args:        []string{"0x01", "0x02"},
		GasLimit:    5000000,
	}
	sdk.Builder.AddAction(contractAction)

	// Add a transfer action
	transferAction := types.WarpTransferAction{
		Type:        types.TransferActionType,
		Label:       "Send EGLD",
		Description: stringPtr("Sends EGLD to a recipient"),
		Address:     stringPtr("erd1..."), // Replace with recipient address
		Value:       stringPtr("0.1"),
	}
	sdk.Builder.AddAction(transferAction)

	// Build the warp
	warp, err := sdk.Builder.Build()
	if err != nil {
		fmt.Println("Error building warp:", err)
		os.Exit(1)
	}

	fmt.Println("Warp built successfully:")
	fmt.Printf("  Name: %s\n", warp.Name)
	fmt.Printf("  Title: %s\n", warp.Title)
	fmt.Printf("  Description: %s\n", *warp.Description)
	fmt.Printf("  Actions: %d\n", len(warp.Actions))

	// Create a transaction for the warp
	tx, err := sdk.Builder.CreateInscriptionTransaction(warp)
	if err != nil {
		fmt.Println("Error creating transaction:", err)
		os.Exit(1)
	}
	fmt.Println("Transaction created:", tx[:64]+"...")

	// Example 2: Create a warp link
	fmt.Println("\nExample 2: Create a warp link")
	
	// Assume we have a hash from a previously created warp
	hash := "your-hash-id"
	url := sdk.Link.Build(types.HashIDType, hash)
	fmt.Println("Warp link:", url)

	// Example 3: Generate a QR code
	fmt.Println("\nExample 3: Generate a QR code")
	qrCode, err := sdk.Link.GenerateQRCode(types.HashIDType, hash, 256)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		os.Exit(1)
	}
	fmt.Printf("QR code generated (%d bytes)\n", len(qrCode))

	// Save the QR code to a file
	err = os.WriteFile("unified-warp-qr.png", qrCode, 0644)
	if err != nil {
		fmt.Println("Error saving QR code:", err)
		os.Exit(1)
	}
	fmt.Println("QR code saved to unified-warp-qr.png")

	// Example 4: Search for warps
	fmt.Println("\nExample 4: Search for warps")
	searchResults, err := sdk.Registry.Search("example")
	if err != nil {
		fmt.Println("Error searching for warps:", err)
		// Don't exit, continue with the example
	} else {
		fmt.Printf("Found %d warps\n", len(searchResults.Hits))
		for i, hit := range searchResults.Hits {
			if i >= 3 {
				fmt.Println("  ...")
				break
			}
			fmt.Printf("  %d. %s - %s\n", i+1, hit.Title, hit.Description)
		}
	}

	// Example 5: Detect a warp from a URL
	fmt.Println("\nExample 5: Detect a warp from a URL")
	warpURL := fmt.Sprintf("https://usewarp.to/to?warp=hash:%s", hash)
	detection, err := sdk.Link.Detect(warpURL)
	if err != nil {
		fmt.Println("Error detecting warp:", err)
		// Don't exit, continue with the example
	} else {
		if detection.Match {
			fmt.Println("Warp detected:")
			fmt.Printf("  URL: %s\n", detection.URL)
			if detection.Warp != nil {
				fmt.Printf("  Name: %s\n", detection.Warp.Name)
				fmt.Printf("  Title: %s\n", detection.Warp.Title)
			}
			if detection.RegistryInfo != nil {
				fmt.Printf("  Creator: %s\n", detection.RegistryInfo.Creator)
				fmt.Printf("  Trust: %s\n", detection.RegistryInfo.Trust)
			}
		} else {
			fmt.Println("No warp detected in URL")
		}
	}

	fmt.Println("\nExample completed successfully!")
}

// Helper function to create a string pointer
func stringPtr(s string) *string {
	return &s
} 