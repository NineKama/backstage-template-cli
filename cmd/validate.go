package cmd

import (
	"fmt"
	"os"

	"backstage-template/internal/model"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a Backstage Component YAML file",
	Long:  `Validates the structure and fields of a Backstage Component YAML file to ensure compliance.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if a file argument is provided
		if len(args) < 1 {
			fmt.Println("Error: Please provide a file path to validate.")
			os.Exit(1)
		}

		// Read the file
		filePath := args[0]
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}

		// Parse the YAML
		var component model.Component
		err = yaml.Unmarshal(data, &component)
		if err != nil {
			fmt.Printf("Error parsing YAML: %v\n", err)
			os.Exit(1)
		}

		// Validate the required fields
		if component.APIVersion == "" || component.Kind == "" {
			fmt.Println("Error: 'apiVersion' and 'kind' fields are required.")
			os.Exit(1)
		}
		if component.Metadata.Name == "" {
			fmt.Println("Error: 'metadata.name' is required.")
			os.Exit(1)
		}
		if component.Metadata.Annotations["github.com/project-slug"] == "" {
			fmt.Println("Error: 'metadata.annotations.github.com/project-slug' is required.")
			os.Exit(1)
		}
		if component.Spec.Type == "" || component.Spec.Owner == "" || component.Spec.Lifecycle == "" {
			fmt.Println("Error: 'spec.type', 'spec.owner', and 'spec.lifecycle' are required.")
			os.Exit(1)
		}

		// If all validations pass
		fmt.Println("Validation successful! The YAML file is valid.")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
