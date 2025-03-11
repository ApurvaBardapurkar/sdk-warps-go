package main

import (
	"fmt"
	"os"

	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/builder"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/link"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
)

func main() {
	// Create a config
	config := types.WarpConfig{
		Env:         types.Mainnet,
		ClientURL:   "https://usewarp.to/to",
		UserAddress: "erd1...", // Replace with your wallet address
	}

	// Example 1: Create a warp link
	fmt.Println("Example 1: Create a warp link")
	warpLink := link.NewWarpLink(config)
	url := warpLink.Build(types.HashIDType, "your-hash-id")
	fmt.Println("Warp link:", url)

	// Example 2: Create a warp
	fmt.Println("\nExample 2: Create a warp")
	warpBuilder := builder.NewWarpBuilder(config)
	warpBuilder.SetName("my-first-warp")
	warpBuilder.SetTitle("My First Warp")
	warpBuilder.SetDescription("This is my first warp created with the Go SDK")

	// Add a contract action
	contractAction := types.WarpContractAction{
		Type:        types.ContractActionType,
		Label:       "Deploy Contract",
		Description: stringPtr("Deploys a new smart contract"),
		Address:     "erd1qqqqqqqqqqqqqpgqt4g39q9vzm80aujs44wmvwn4kfn6czmtcrpslzjs4e",
		Func:        stringPtr("deployContract"),
		Args:        []string{"0x01", "0x02"},
		GasLimit:    10000000,
	}
	warpBuilder.AddAction(contractAction)

	// Add a link action
	linkAction := types.WarpLinkAction{
		Type:        types.LinkActionType,
		Label:       "Visit Website",
		Description: stringPtr("Visit our website"),
		URL:         "https://example.com",
	}
	warpBuilder.AddAction(linkAction)

	// Build the warp
	warp, err := warpBuilder.Build()
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
	tx, err := warpBuilder.CreateInscriptionTransaction(warp)
	if err != nil {
		fmt.Println("Error creating transaction:", err)
		os.Exit(1)
	}
	fmt.Println("Transaction created:", tx[:64]+"...")

	// Example 3: Generate a QR code
	fmt.Println("\nExample 3: Generate a QR code")
	qrCode, err := warpLink.GenerateQRCode(types.HashIDType, "your-hash-id", 256)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		os.Exit(1)
	}
	fmt.Printf("QR code generated (%d bytes)\n", len(qrCode))

	// Save the QR code to a file
	err = os.WriteFile("warp-qr.png", qrCode, 0644)
	if err != nil {
		fmt.Println("Error saving QR code:", err)
		os.Exit(1)
	}
	fmt.Println("QR code saved to warp-qr.png")
}

// Helper function to create a string pointer
func stringPtr(s string) *string {
	return &s
} 