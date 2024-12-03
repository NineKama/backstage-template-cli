package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"backstage-template/internal/model"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Backstage component YAML file",
	Long:  `Prompts the user for input and generates a Backstage component YAML file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the Component struct
		component := model.Component{
			APIVersion: "backstage.io/v1alpha1",
			Kind:       "Component",
		}

		// Prompt the user for input
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter component name: ")
		component.Metadata.Name = readInput(reader)

		fmt.Print("Enter description: ")
		component.Metadata.Description = readInput(reader)

		fmt.Print("Enter GitHub project slug (e.g., user/repo): ")
		component.Metadata.Annotations = map[string]string{
			"github.com/project-slug": readInput(reader),
		}

		fmt.Print("Enter type (e.g., service): ")
		component.Spec.Type = readInput(reader)

		fmt.Print("Enter owner: ")
		component.Spec.Owner = readInput(reader)

		fmt.Print("Enter lifecycle (e.g., production): ")
		component.Spec.Lifecycle = readInput(reader)

		// Convert struct to YAML
		output, err := yaml.Marshal(&component)
		if err != nil {
			fmt.Println("Error generating YAML:", err)
			return
		}

		// Write YAML to file
		fileName := component.Metadata.Name + ".yaml"
		err = os.WriteFile(fileName, output, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Printf("YAML file created: %s\n", fileName)
	},
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
