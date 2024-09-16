package module

import (
	"fmt"
	"log"

	"github.com/trinitytechnology/ebrick-cli/internal/constants"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

type ModuleConfig struct {
	Name      string `yaml:"name"`
	External  bool   `yaml:"external"`
	Rest      bool   `yaml:"rest"`
	Graphql   bool   `yaml:"graphql"`
	Grpc      bool   `yaml:"grpc"`
	Messaging bool   `yaml:"messaging"`
	Auth      bool   `yaml:"auth"`
}

func NewModule() {

	// Check .ebrick.yaml file exists
	if !utils.FileExists(constants.AppManifestFile) {
		log.Fatalf("No %s file found. Please run 'ebrick new app' to create a new application.", constants.AppManifestFile)
	}

	moduleConfig := newModuleCommandPrompts()

	// Create Module configuration file
	err := utils.WriteYamlFile(fmt.Sprintf("%s/%s.yaml", constants.ModuleManifestDir, moduleConfig.Name), moduleConfig)

	if err != nil {
		log.Fatalf("Error writing module configuration file: %s", err)
	}

	// Execute post generation tasks
	PostGenerated()
}

func newModuleCommandPrompts() ModuleConfig {
	var moduleConfig ModuleConfig
	moduleName := utils.GetUserInput("Enter the name of the module: ", true, "Module name is required.")
	moduleConfigFile := fmt.Sprintf("%s/%s.yaml", constants.ModuleManifestDir, moduleName)

	// Check if module already exists and ask if user wants to regenerate
	if utils.FileExists(moduleConfigFile) {
		regenerate := utils.GetYesOrNoInput("Module already exists. Do you want to regenerate the configuration?", true)
		if regenerate {
			// Module already exists, read the configuration file
			moduleConfig, err := utils.ReadYamlFile[ModuleConfig](moduleConfigFile)
			if err != nil {
				log.Fatalf("Error reading module configuration file: %s", err)
			}
			return moduleConfig
		}
	}

	// Get module configuration from user
	external := utils.GetYesOrNoInput("Is this an external module?", false)
	rest := utils.GetYesOrNoInput("Do you want to enable REST API?", false)
	graphql := utils.GetYesOrNoInput("Do you want to enable GraphQL API?", false)
	grpc := utils.GetYesOrNoInput("Do you want to enable gRPC API?", false)
	messaging := utils.GetYesOrNoInput("Do you want to enable Messaging?", false)
	auth := utils.GetYesOrNoInput("Do you want to enable Authentication?", false)
	moduleConfig = ModuleConfig{
		Name:      moduleName,
		External:  external,
		Rest:      rest,
		Graphql:   graphql,
		Grpc:      grpc,
		Messaging: messaging,
		Auth:      auth,
	}
	return moduleConfig
}

func PostGenerated() {
	fmt.Println("Post generated tasks")
}
