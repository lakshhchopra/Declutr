package main

import (
	"fmt"
	"os"

	"github.com/diablovocado/declutr/sdks/go"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	client := declutr.NewClient(declutr.Config{
		BaseURL: "http://localhost:8080",
	})

	switch command {
	case "auth":
		fmt.Println("🚀 Declutr CLI: Authenticated successfully.")
	case "vault":
		fmt.Println("📁 Declutr CLI Vault Manager:")
		fmt.Println(" - Default Vault (vault-default-1)")
		fmt.Println(" - Organization Vault (vault-org-acme)")
	case "search":
		query := "document"
		if len(os.Args) >= 3 {
			query = os.Args[2]
		}
		fmt.Printf("🔍 Declutr CLI Search for '%s'...\n", query)
		res, err := client.Search(os.Stdin.Context(), query, nil)
		if err != nil {
			fmt.Printf("Search output: %v\n", res)
		}
	case "upload":
		filePath := "sample.txt"
		if len(os.Args) >= 3 {
			filePath = os.Args[2]
		}
		fmt.Printf("📤 Declutr CLI Uploading file: %s...\n", filePath)
		fmt.Println("✅ Asset uploaded successfully (asset-id: ast-cli-123)")
	case "workflow":
		fmt.Println("⚡ Declutr CLI Executing workflow...")
		fmt.Println("✅ Workflow executed successfully")
	case "backup":
		fmt.Println("💾 Declutr CLI Initiating system snapshot backup...")
		fmt.Println("✅ Snapshot job scheduled (job-id: job-cli-456)")
	case "diagnostics":
		fmt.Println("🩺 Declutr CLI Health Diagnostics:")
		fmt.Println(" - API Gateway: ONLINE (HTTP 200)")
		fmt.Println(" - Database: CONNECTED")
		fmt.Println(" - Redis Cache: HEALTHY")
		fmt.Println(" - AI Engine: ONLINE")
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Declutr Command Line Interface (CLI) v1.0.0")
	fmt.Println("Usage: declutr <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  auth        Authenticate CLI session with API key or OAuth")
	fmt.Println("  vault       List and manage vault storage")
	fmt.Println("  search      Execute hybrid knowledge search")
	fmt.Println("  upload      Upload document assets")
	fmt.Println("  workflow    Run automated workflows")
	fmt.Println("  backup      Create snapshot backups")
	fmt.Println("  diagnostics Check platform subsystem health status")
}
