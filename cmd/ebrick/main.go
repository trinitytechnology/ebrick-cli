package main

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
	ebrickcli "github.com/trinitytechnology/ebrick-cli"
	"github.com/trinitytechnology/ebrick-cli/internal/app"
)

//go:embed banner.txt
var banner string

func main() {
	// Print the banner with colors
	fmt.Println(banner)

	var rootCmd = &cobra.Command{
		Use: "ebrick",
	}

	rootCmd.AddCommand(createVersionCommand())
	rootCmd.AddCommand(createAppCommands())
	rootCmd.AddCommand(createRunCommand())
	rootCmd.Execute()
}

// add version command
func createVersionCommand() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:     "version",
		Short:   "Print the version number of ebrick",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("eBrick Cli: %s, eBrick Framework: %s \n", ebrickcli.Version, ebrickcli.FrameworkVersion)
		},
	}
	return versionCmd
}

func createAppCommands() *cobra.Command {
	var newCmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new ebrick application, module or service..",
	}

	var newAppCmd = &cobra.Command{
		Use:   "app",
		Short: "Create a new ebrick application",
		Run: func(cmd *cobra.Command, args []string) {
			app.NewApp()
		},
	}
	newCmd.AddCommand(newAppCmd)

	return newCmd
}

func createRunCommand() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the ebrick application",
		Run: func(cmd *cobra.Command, args []string) {
			app.RunApp()
		},
	}
	return runCmd
}
