package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yourusername/airfoil/cmd/project"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "airfoil",
		Short: "Airfoil is a CLI tool for managing RunPod projects",
		Long: `Airfoil is a command-line interface for developing and deploying projects on RunPod's infrastructure.
It provides a seamless workflow for creating, developing, and deploying AI and machine learning projects.`,
		Version: "1.0.0",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.airfoil.yaml)")

	project.InitializeCommands(rootCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of Airfoil",
		Long:  `All software has versions. This is Airfoil's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Airfoil v%s\n", rootCmd.Version)
		},
	})
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".airfoil")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
