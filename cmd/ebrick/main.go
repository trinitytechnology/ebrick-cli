package main

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trinitytechnology/ebrick-cli/internal/app"
	"github.com/trinitytechnology/ebrick-cli/internal/module"
)

var version = "development"
var frameworkVersion = "v0.3.5"

//go:embed banner.txt
var banner string

func main() {
	// Print the banner with colors
	fmt.Println(banner)

	var rootCmd = &cobra.Command{
		Use: "ebrick",
	}

	rootCmd.AddCommand(versionCommand())
	rootCmd.AddCommand(newCommand())
	rootCmd.AddCommand(runCommand())
	rootCmd.Execute()
}

// add version command
func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Print the version number of ebrick",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("eBrick Cli: %s, eBrick Framework: %s \n", version, frameworkVersion)
		},
	}
}

func newCommand() *cobra.Command {
	var newCmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new ebrick application, module or service..",
	}
	newCmd.AddCommand(newAppCommand())
	newCmd.AddCommand(newModuleCommand())

	return newCmd
}

func newAppCommand() *cobra.Command {

	return &cobra.Command{
		Use:   "app",
		Short: "Create a new ebrick application",
		Run: func(cmd *cobra.Command, args []string) {
			app.NewApp(frameworkVersion)
		},
	}
}

func newModuleCommand() *cobra.Command {

	return &cobra.Command{
		Use:   "module",
		Short: "Create a new ebrick module",
		Run: func(cmd *cobra.Command, args []string) {
			module.NewModule()
		},
	}
}

func runCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run the ebrick application",
		Run: func(cmd *cobra.Command, args []string) {
			app.RunApp()
		},
	}
}
