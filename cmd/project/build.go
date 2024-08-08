package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	outputPath string
	tag        string
)

var BuildProjectCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Dockerfile for current project",
	Long: `Builds a local Dockerfile for the project in the current folder.
You can use this Dockerfile to build an image and deploy it to any API server.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Building Dockerfile...")
		// Implement the logic for building a Dockerfile here

		// Example of how you might log during the build process:
		// if err := buildDockerfile(); err != nil {
		//     logger.Error("Failed to build Dockerfile", zap.Error(err))
		//     return
		// }

	},
}

func init() {
	BuildProjectCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path for the Dockerfile (default is ./Dockerfile)")
	BuildProjectCmd.Flags().StringVarP(&tag, "tag", "t", "", "Suggest a tag for the Docker image")
	BuildProjectCmd.Flags().BoolVar(&includeEnvInDockerfile, "include-env", false, "Incorporate environment variables defined in runpod.toml into the generated Dockerfile")

}
