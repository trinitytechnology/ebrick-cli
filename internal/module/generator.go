package module

import (
	_ "embed"
	"fmt"

	"github.com/trinitytechnology/ebrick-cli/internal/app"
	"github.com/trinitytechnology/ebrick-cli/internal/templates"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

var files = map[string]string{}

type ModuleGenerator struct {
	appConfig    *app.AppConfig
	moduleConfig *ModuleConfig
	moduleDir    string
}

func NewModuleGenerator(appConfig *app.AppConfig, moduleConfig *ModuleConfig) *ModuleGenerator {
	return &ModuleGenerator{
		moduleConfig: moduleConfig,
		moduleDir:    MODULE_INTERNAL_DIR + "/" + moduleConfig.Package,
		appConfig:    appConfig,
	}
}

func (m ModuleGenerator) Generate() {
	fmt.Println("Creating a new module with the name:", m.moduleConfig.Name)

	// Create module directory
	utils.CreateFolder(MODULE_INTERNAL_DIR + "/" + m.moduleConfig.Package)

	// Generate module files
	m.generateModuleFiles()

	// execute post generation tasks
	m.postGenerated()
}

func (m ModuleGenerator) generateModuleFiles() {
	files = make(map[string]string)
	files[m.moduleDir+"/"+m.moduleConfig.Package+".go"] = templates.ModuleTemplate

	for file, content := range files {
		utils.GenerateFileFromTemplate(file, m.moduleConfig, content)
		fmt.Println("Generated", file, "successfully.")
	}
}

func (m ModuleGenerator) postGenerated() {
	fmt.Println("Post generated tasks")
}
