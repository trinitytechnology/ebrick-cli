package app

import (
	_ "embed"
	"fmt"

	"github.com/trinitytechnology/ebrick-cli/internal/model"
	"github.com/trinitytechnology/ebrick-cli/internal/templates"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

type AppGenerator struct {
	ebrickApp *model.EBrickApp
	files     map[string]string
}

func NewAppGenerator(ebrickApp *model.EBrickApp) *AppGenerator {

	files := map[string]string{
		templates.FILE_APPLICATION_YAML: templates.ApplicationTemplate,
		templates.FILE_MAIN:             templates.MainTemplate,
		templates.FILE_DOCKER_COMPOSE:   templates.DockerComposeTemplate,
		templates.FILE_GO_MOD:           templates.GoModTemplate,
		templates.FILE_README:           templates.ReadmeTemplate,
		templates.FILE_DOCKERFILE:       templates.DockerfileTemplate,
		templates.FILE_DOCKER_APP:       templates.DockerAppTemplate,
	}
	if ebrickApp.Observability {
		files[templates.FILE_GRAFANA_PROMETHEUS] = templates.GrafanaPrometheusTemplate
		files[templates.FILE_GRAFANA_DATASOURCE] = templates.GrafanaDatasourceTemplate
	}
	return &AppGenerator{
		ebrickApp: ebrickApp,
		files:     files,
	}
}

func (m AppGenerator) Generate() {
	// Create the necessary folders
	m.createFolders()

	// Generate the application.yaml file
	m.generateFiles()

	m.postGenerated()

}

func (m AppGenerator) createFolders() {
	fmt.Println("Creating the necessary folders...")
	utils.CreateFolder("cmd")
	utils.CreateFolder("modules")
	utils.CreateFolder("internal")
	utils.CreateFolder("pkg")
}

func (m AppGenerator) generateFiles() {
	for file, template := range m.files {
		utils.GenerateFileFromTemplate(file, m.ebrickApp, template)
		fmt.Println("Generated", file, "successfully.")
	}
}

func (m AppGenerator) postGenerated() {
	fmt.Println("Running post generation tasks...")

	// Run go mod tidy
	utils.ExecCommand("go", "mod", "tidy")

}
