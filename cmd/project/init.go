package project

import (
	"github.com/spf13/cobra"
)

// InitializeCommands adds all project-related commands to the root command
func InitializeCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(NewProjectCmd)
	rootCmd.AddCommand(StartProjectCmd)
	rootCmd.AddCommand(DeployProjectCmd)
	rootCmd.AddCommand(BuildProjectCmd)
}
