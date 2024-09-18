package app

import (
	"fmt"
	"os"

	"github.com/trinitytechnology/ebrick-cli/internal/constants"
	"github.com/trinitytechnology/ebrick-cli/internal/model"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

// NewApp creates a new eBrick application
func NewApp(frameworkVersion string) {

	var ebrickApp model.EBrickApp

	// Check .ebrick.yaml file exists
	if !utils.FileExists(constants.AppManifestFile) {
		ebrickApp = NewApplicationCommandPrompts(frameworkVersion)
	} else {
		overwrite := utils.GetYesOrNoInput("Overwrite existing configuration?", true)
		if !overwrite {
			ebrickApp = NewApplicationCommandPrompts(frameworkVersion)
		}
	}

	// Read .ebrick.yaml
	ebrickApp, err := utils.ReadYamlFile[model.EBrickApp](constants.AppManifestFile)
	if err != nil {
		fmt.Println("Error reading .ebrick.yaml:", err)
		return
	}

	fmt.Println("Creating a new eBrick application with the name:", ebrickApp.Name)
	generator := NewAppGenerator(&ebrickApp)
	generator.Generate()

	fmt.Println("Application created successfully.")

}

func NewApplicationCommandPrompts(frameworkVersion string) model.EBrickApp {

	appName := utils.GetUserInput("Enter the name of the application: ", true, "Application name is required.")
	packageName := utils.GetUserInput("Enter the application package: ", true, "Package name is required.")
	modulesInput := utils.GetUserInput("Enter the application modules (comma-separated, no spaces): ", false, "")
	extModules := utils.ProcessSlicesInput(modulesInput)

	database := utils.GetYesOrNoInput("Do you need a database?", true)
	cache := utils.GetYesOrNoInput("Do you need a cache?", false)
	messaging := utils.GetYesOrNoInput("Do you need messaging?", false)
	observability := utils.GetYesOrNoInput("Do you need observability?", false)

	ebrickApp := model.EBrickApp{
		Name:            appName,
		Package:         packageName,
		ExternalModules: extModules,
		Database:        database,
		Observability:   observability,
		Cache:           cache,
		Messaging:       messaging,
		Version:         frameworkVersion,
		InternalModules: []model.Module{},
	}
	err := utils.WriteYamlFile(constants.AppManifestFile, ebrickApp)
	if err != nil {
		os.Exit(1)
	}
	return ebrickApp
}

func RunApp() {
	// Run go mod tidy
	utils.ExecCommand("go", "mod", "tidy")

	// Run go mod tidy
	utils.ExecCommand("go", "run", "cmd/main.go")
}
