package module

import (
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"github.com/trinitytechnology/ebrick-cli/internal/constants"
	"github.com/trinitytechnology/ebrick-cli/internal/model"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

func NewModule() {

	// Check .ebrick.yaml file exists
	if !utils.FileExists(constants.AppManifestFile) {
		log.Fatalf("No %s file found. Please run 'ebrick new app' to create a new application.", constants.AppManifestFile)
	}

	// Read .ebrick.yaml
	ebrickApp, err := utils.ReadYamlFile[*model.EBrickApp](constants.AppManifestFile)
	if err != nil {
		log.Fatalf("Error reading %s file: %s", constants.AppManifestFile, err)
	}
	// Get module configuration from user
	module := newModuleCommandPrompts(ebrickApp)

	generator := NewModuleGenerator(ebrickApp, module)
	generator.Generate()
}

func findModule(name string, ebrickApp *model.EBrickApp) *model.Module {
	for _, module := range ebrickApp.InternalModules {
		if strings.EqualFold(module.Name, name) {
			return &module
		}
	}
	return nil
}

// NewModuleCommandPrompts prompts the user for module configuration
func newModuleCommandPrompts(ebrickApp *model.EBrickApp) *model.Module {
	var module *model.Module
	moduleName := utils.GetUserInput("Enter the name of the module: ", true, "Module name is required.")

	module = findModule(moduleName, ebrickApp)
	// Check if module already exists and ask if user wants to regenerate
	if module != nil {
		if regenerate := utils.GetYesOrNoInput("Module already exists. Do you want to regenerate the configuration?", true); regenerate {
			return module
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
	if ebrickApp.Messaging {
		messaging = utils.GetYesOrNoInput("Do you want to use Messaging?", false)
	}

	auth := utils.GetYesOrNoInput("Do you want to enable Authentication?", false)

	// Check if module already exists then update the configuration
	if module != nil {
		module.Version = version
		module.Description = description
		module.External = external
		module.Rest = rest
		module.Graphql = graphql
		module.Grpc = grpc
		module.Messaging = messaging
		module.Auth = auth
		return module
	} else {
		newModule := model.Module{
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

		ebrickApp.InternalModules = append(ebrickApp.InternalModules, newModule)
		return &newModule
	}
}
