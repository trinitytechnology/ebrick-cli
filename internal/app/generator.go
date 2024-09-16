package app

import (
	_ "embed"
	"fmt"

	"github.com/trinitytechnology/ebrick-cli/internal/templates"
	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

type AppGenerator struct {
	appConfig *AppConfig
	files     map[string]string
}

func NewAppGenerator(appConfig *AppConfig) *AppGenerator {
	files := make(map[string]string)
	files["application.yaml"] = templates.ApplicationTemplate
	files["cmd/main.go"] = templates.MainTemplate
	files["docker-compose.yml"] = templates.DockerComposeTemplate
	files["go.mod"] = templates.GoModTemplate
	files["README.md"] = templates.ReadmeTemplate
	files["Dockerfile"] = templates.DockerfileTemplate
	if appConfig.Observability {
		files["observability/prometheus/prometheus.yml"] = templates.GrafanaPrometheusTemplate
		files["observability/grafana/datasource.yml"] = templates.GrafanaDatasourceTemplate
	}

	return &AppGenerator{
		appConfig: appConfig,
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
		utils.GenerateFileFromTemplate(file, m.appConfig, template)
		fmt.Println("Generated", file, "successfully.")
	}
}

func (m AppGenerator) postGenerated() {
	fmt.Println("Running post generation tasks...")

	// Run go mod tidy
	utils.ExecCommand("go", "mod", "tidy")

}
