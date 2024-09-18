package module

import (
	_ "embed"
	"fmt"

	"github.com/trinitytechnology/ebrick-cli/internal/constants"
	"github.com/trinitytechnology/ebrick-cli/internal/model"
	"github.com/trinitytechnology/ebrick-cli/internal/templates"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

type ModuleGenerator struct {
	ebrickApp    *model.EBrickApp
	moduleConfig *model.Module
	moduleDir    string
	files        map[string]string
}

func NewModuleGenerator(ebrickApp *model.EBrickApp, moduleConfig *model.Module) *ModuleGenerator {
	moduleDir := MODULE_INTERNAL_DIR + "/" + moduleConfig.Package

	files := make(map[string]string)
	files[moduleDir+"/"+moduleConfig.Package+".go"] = templates.ModuleTemplate

	return &ModuleGenerator{
		moduleConfig: moduleConfig,
		moduleDir:    MODULE_INTERNAL_DIR + "/" + moduleConfig.Package,
		ebrickApp:    ebrickApp,
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

	// Add module to app manifest
	utils.WriteYamlFile(constants.AppManifestFile, m.ebrickApp)
	// Add module to main.go
	utils.GenerateFileFromTemplate(templates.FILE_MAIN, m.ebrickApp, templates.MainTemplate)

}

func (m ModuleGenerator) postGenerated() {
	fmt.Println("Post generated tasks")
}
