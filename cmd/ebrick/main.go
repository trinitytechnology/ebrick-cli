package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/trinitytechnology/ebrick-cli/internal/app"
	"github.com/trinitytechnology/ebrick-cli/internal/module"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
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
	rootCmd.AddCommand(buildApp())
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

func buildApp() *cobra.Command {
	var ldflags string
	var output string

	cmd := &cobra.Command{
		Use:   "build",
		Short: "build the ebrick application, can pass flags go build, such as -ldflags \"-X main.version=1234\"",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Building the application...")

			// Run go mod tidy
			if err := utils.ExecCommand("go", "mod", "tidy"); err != nil {
				log.Fatalf("Error running go mod tidy: %v", err)
			}

			// Prepare arguments for go build, including ldflags if provided
			buildArgs := []string{"build", "-o", output}
			if ldflags != "" {
				buildArgs = append(buildArgs, "-ldflags", ldflags)
			}
			buildArgs = append(buildArgs, "cmd/main.go")

			// Run go build
			if err := utils.ExecCommand("go", buildArgs...); err != nil {
				log.Fatalf("Error building the application: %v", err)
			}

			log.Println("Application built successfully!")
		},
	}

	// Add ldflags as a command-line flag
	cmd.Flags().StringVarP(&ldflags, "ldflags", "l", "", "Flags to pass to go build, such as -ldflags \"-X main.version=1234\"")
	cmd.Flags().StringVarP(&output, "output", "o", "app", "Output file name")
	return cmd
}
