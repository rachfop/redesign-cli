package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
	"github.com/yourusername/airfoil/cmd/project"
)

var (
	cfgFile      string
	generateDocs bool
	versionFlag  bool
	rootCmd      = &cobra.Command{
		Use:   "airfoil",
		Short: "Airfoil is a CLI tool for managing RunPod projects",
		Long: `Airfoil is a command-line interface for developing and deploying projects on RunPod's infrastructure.
It provides commands for creating new projects, starting development sessions, deploying projects, and building Dockerfiles.`,
		Version: "1.0.0",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.airfoil.yaml)")
	rootCmd.Flags().BoolVar(&generateDocs, "generate-docs", false, "Generate documentation")
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Print the version number of Airfoil")

	rootCmd.AddCommand(project.NewProjectCmd)
	rootCmd.AddCommand(project.StartProjectCmd)
	rootCmd.AddCommand(project.DeployProjectCmd)
	rootCmd.AddCommand(project.BuildProjectCmd)
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
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func main() {

	if versionFlag {
		fmt.Println(rootCmd.Version)
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateDocumentation() error {
	docsDir := "./docs"
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		return fmt.Errorf("creating docs directory: %w", err)
	}

	if err := doc.GenMarkdownTree(rootCmd, docsDir); err != nil {
		return fmt.Errorf("generating markdown documentation: %w", err)
	}

	manDir := "./man"
	if err := os.MkdirAll(manDir, 0755); err != nil {
		return fmt.Errorf("creating man directory: %w", err)
	}

	header := &doc.GenManHeader{
		Title:   "AIRFOIL",
		Section: "1",
	}
	if err := doc.GenManTree(rootCmd, header, manDir); err != nil {
		return fmt.Errorf("generating man pages: %w", err)
	}

	if err := generateShellCompletions(); err != nil {
		return fmt.Errorf("generating shell completions: %w", err)
	}

	return nil
}

func generateShellCompletions() error {
	completionDir := "./completions"
	if err := os.MkdirAll(completionDir, 0755); err != nil {
		return fmt.Errorf("creating completions directory: %w", err)
	}

	shells := []struct {
		name string
		fn   func(string) error
	}{
		{"bash", rootCmd.GenBashCompletionFile},
		{"zsh", rootCmd.GenZshCompletionFile},
		{"fish", func(path string) error {
			return rootCmd.GenFishCompletionFile(path, true)
		}},
		{"powershell", rootCmd.GenPowerShellCompletionFile},
	}

	for _, shell := range shells {
		file := filepath.Join(completionDir, "airfoil."+shell.name)
		if err := shell.fn(file); err != nil {
			return fmt.Errorf("generating %s completion: %w", shell.name, err)
		}
	}

	return nil
}
