package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"devstation-cli/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "devstation",
	Short: "A CLI tool to set up Python and C development environment on Windows",
	Long: `DevStation CLI helps you quickly set up a complete development environment
for Python and C programming on Windows machines. It automates the installation
of essential tools, compilers, and development dependencies.`,
}

func main() {
	// Initialize commands
	cmd.InitCommands(rootCmd)
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
