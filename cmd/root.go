package cmd

import (
	"fmt"
	"os"

	"net-tools/internal/scanner"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "net-tools",
	Short: "A network utility toolkit",
	Long: `Net-tools is a CLI toolkit for network utilities.
Select from various network tools to perform network analysis and scanning.`,
	Run: func(cmd *cobra.Command, args []string) {
		showToolSelection()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func showToolSelection() {
	fmt.Println("=== Net-Tools ===")
	fmt.Println("Select a tool:")
	fmt.Println("1. Port Scanner")
	fmt.Println("0. Exit")
	fmt.Print("\nEnter your choice: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("\nStarting Port Scanner...")
		runScanner()
	case 0:
		fmt.Println("Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
		showToolSelection()
	}
}

func runScanner() {
	s := scanner.NewScanner()
	s.RunInteractiveMode()

	// After scanning, show the tool selection again
	fmt.Println("\nPress Enter to return to main menu...")
	fmt.Scanln()
	showToolSelection()
}
