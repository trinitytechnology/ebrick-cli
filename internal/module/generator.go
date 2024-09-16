package module

import (
	_ "embed"
	"fmt"

	"github.com/trinitytechnology/ebrick-cli/internal/app"
	"github.com/trinitytechnology/ebrick-cli/internal/templates"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

type ModuleGenerator struct {
	appConfig    *app.AppConfig
	moduleConfig *ModuleConfig
	moduleDir    string
	files        map[string]string
}

func NewModuleGenerator(appConfig *app.AppConfig, moduleConfig *ModuleConfig) *ModuleGenerator {
	moduleDir := MODULE_INTERNAL_DIR + "/" + moduleConfig.Package

	files := make(map[string]string)
	files[moduleDir+"/"+moduleConfig.Package+".go"] = templates.ModuleTemplate

	return &ModuleGenerator{
		moduleConfig: moduleConfig,
		moduleDir:    MODULE_INTERNAL_DIR + "/" + moduleConfig.Package,
		appConfig:    appConfig,
		files:        files,
	}
}

func (m ModuleGenerator) Generate() {
	fmt.Println("Creating a new module with the name:", m.moduleConfig.Name)

	// Create module directory
	utils.CreateFolder(m.moduleDir)

	// Generate module files
	m.generateModuleFiles()

	// execute post generation tasks
	m.postGenerated()
}

func (m ModuleGenerator) generateModuleFiles() {

	for file, content := range m.files {
		utils.GenerateFileFromTemplate(file, m.moduleConfig, content)
		fmt.Println("Generated", file, "successfully.")
	}
}

func (m ModuleGenerator) postGenerated() {
	fmt.Println("Post generated tasks")
}
