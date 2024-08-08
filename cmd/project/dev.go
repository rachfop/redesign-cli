package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StartProjectCmd = &cobra.Command{
	Use:     "dev",
	Aliases: []string{"start"},
	Short:   "Start a development session for the current project",
	Long:    "This command establishes a connection between your local development environment and your RunPod project environment, allowing for real-time synchronization of changes.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting a development session...")
		// Implement the logic for starting a project here
	},
}

func init() {
	StartProjectCmd.Flags().BoolVar(&setDefaultNetworkVolume, "select-volume", false, "Choose a new default network volume for the project")
	StartProjectCmd.Flags().BoolVar(&showPrefixInPodLogs, "prefix-pod-logs", true, "Include the Pod ID as a prefix in log messages from the project Pod")
}
