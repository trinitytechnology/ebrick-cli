package module

import (
	_ "embed"
	"fmt"

	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

var files = map[string]string{}

type ModuleGenerator struct {
	moduleConfig *ModuleConfig
	moduleDir    string
}

func NewModuleGenerator(moduleConfig *ModuleConfig) *ModuleGenerator {
	return &ModuleGenerator{
		moduleConfig: moduleConfig,
		moduleDir:    MODULE_INTERNAL_DIR + "/" + moduleConfig.Package,
	}
}

func (m ModuleGenerator) Generate() {
	fmt.Println("Creating a new module with the name:", m.moduleConfig.Name)

	// Create module directory
	utils.CreateFolder(MODULE_INTERNAL_DIR + "/" + m.moduleConfig.Package)

	// Generate module files
	m.generateFiles()

	// execute post generation tasks
	m.postGenerated()
}

//go:embed templates/module.go.tmpl
var moduleTemplate string

func (m ModuleGenerator) generateFiles() {
	files = make(map[string]string)
	files[m.moduleDir+"/"+m.moduleConfig.Package+".go"] = moduleTemplate

	for file, content := range files {
		utils.GenerateFileFromTemplate(file, m.moduleConfig, content)
		fmt.Println("Generated", file, "successfully.")
	}
}

func (m ModuleGenerator) postGenerated() {
	fmt.Println("Post generated tasks")
}
