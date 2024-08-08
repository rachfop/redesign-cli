package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeployProjectCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys your project as an endpoint",
	Long:  "Deploys a serverless endpoint for the RunPod project in the current folder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deploying project...")
		// Implement the logic for deploying a project here
	},
}
