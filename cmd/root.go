package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "backstage-cli",
	Short: "A CLI to manage Backstage components",
	Long:  "A Command Line Interface to generate and manage Backstage YAML files.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add the generate command
	rootCmd.AddCommand(generateCmd)
}
