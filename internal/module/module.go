package module

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"github.com/trinitytechnology/ebrick-cli/internal/app"
	"github.com/trinitytechnology/ebrick-cli/internal/constants"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

type ModuleConfig struct {
	Id          string `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Package     string `yaml:"package"`
	External    bool   `yaml:"external"`
	Rest        bool   `yaml:"rest"`
	Graphql     bool   `yaml:"graphql"`
	Grpc        bool   `yaml:"grpc"`
	Messaging   bool   `yaml:"messaging"`
	Auth        bool   `yaml:"auth"`
}

func NewModule() {

	// Check .ebrick.yaml file exists
	if !utils.FileExists(constants.AppManifestFile) {
		log.Fatalf("No %s file found. Please run 'ebrick new app' to create a new application.", constants.AppManifestFile)
	}

	// Read .ebrick.yaml
	appConfig, err := utils.ReadYamlFile[app.AppConfig](constants.AppManifestFile)
	if err != nil {
		log.Fatalf("Error reading %s file: %s", constants.AppManifestFile, err)
	}

	// Get module configuration from user
	moduleConfig := newModuleCommandPrompts(appConfig)

	// Create Module configuration file
	err = utils.WriteYamlFile(fmt.Sprintf("%s/%s.yaml", constants.ModuleManifestDir, moduleConfig.Name), moduleConfig)
	if err != nil {
		log.Fatalf("Error writing module configuration file: %s", err)
	}

	generator := NewModuleGenerator(&moduleConfig)
	generator.Generate()

}

// NewModuleCommandPrompts prompts the user for module configuration
func newModuleCommandPrompts(appConfig app.AppConfig) ModuleConfig {
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

	version := utils.GetUserInputWithValidation("Enter the version of the module(eg v1.0.0): ", true, "Version is required.", utils.IsValidVersion, "Invalid version format. Please use semantic versioning format.")

	description := utils.GetUserInput("Enter the description of the module: ", false, "")
	external := utils.GetYesOrNoInput("Is this an external module?", false)
	rest := utils.GetYesOrNoInput("Do you want to use REST API?", false)
	graphql := utils.GetYesOrNoInput("Do you want to use GraphQL API?", false)
	grpc := utils.GetYesOrNoInput("Do you want to use gRPC API?", false)

	// check if message enabled then ask for message usage
	var messaging bool
	if appConfig.Messaging {
		messaging = utils.GetYesOrNoInput("Do you want to use Messaging?", false)
	}

	auth := utils.GetYesOrNoInput("Do you want to enable Authentication?", false)

	moduleConfig = ModuleConfig{
		Id:          uuid.New().String(),
		Name:        strcase.ToCamel(moduleName),
		Package:     strcase.ToSnake(moduleName),
		Version:     version,
		Description: description,
		External:    external,
		Rest:        rest,
		Graphql:     graphql,
		Grpc:        grpc,
		Messaging:   messaging,
		Auth:        auth,
	}
	return moduleConfig
}
