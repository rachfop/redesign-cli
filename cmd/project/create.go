package project

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/spf13/cobra"
)

var NewProjectCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Creates a new project",
	Long: `Creates a new RunPod project folder on your local machine.
This command will guide you through the process of setting up a new project,
including selecting a starter template, choosing CUDA and Python versions,
and configuring other project settings.`,
	Example: `  airfoil create --name my-project
  airfoil create --name my-llm-project --type LLM --model gpt2`,
	Run: func(cmd *cobra.Command, args []string) {
		if projectName == "" || modelType == "" || modelName == "" {

			m := initialModel()
			p := tea.NewProgram(m)
			if _, err := p.Run(); err != nil {
				fmt.Println("Failed to create project structure")
				os.Exit(1)
			}
		} else {
			createProjectStructure(projectName, modelType, modelName, cudaVersion, pythonVersion)
			fmt.Println("Project created successfully")

		}

	},
}

func init() {
	NewProjectCmd.Flags().StringVarP(&projectName, "name", "n", "hello-world", "Set the project name, a directory with this name will be created in the current path")
	NewProjectCmd.Flags().StringVarP(&modelType, "type", "t", "", "Specify the model type for the project")
	NewProjectCmd.Flags().StringVarP(&modelName, "model", "m", "", "Specify the Hugging Face model name for the project")
	NewProjectCmd.Flags().StringVarP(&cudaVersion, "cuda", "c", "12.5", "Specify the CUDA version for the project")
	NewProjectCmd.Flags().StringVarP(&pythonVersion, "python", "p", "3.10", "Specify the Python version for the project")
	NewProjectCmd.Flags().BoolVarP(&initCurrentDir, "init", "i", false, "Initialize the project in the current directory instead of creating a new one")
}
