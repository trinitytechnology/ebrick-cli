package app

import (
	_ "embed"
	"fmt"

	"github.com/trinitytechnology/ebrick-cli/pkg/utils"
)

type AppGenerator struct {
	appConfig *AppConfig
}

func NewAppGenerator(appConfig *AppConfig) *AppGenerator {
	return &AppGenerator{
		appConfig: appConfig,
	}
}

//go:embed templates/application.yaml.tmpl
var applicationTemplate string

//go:embed templates/main.go.tmpl
var mainTemplate string

//go:embed templates/docker-compose.yml.tmpl
var dockerComposeTemplate string

//go:embed templates/go.mod.tmpl
var goModTemplate string

//go:embed templates/README.md.tmpl
var readmeTemplate string

//go:embed templates/observability/prometheus/prometheus.yml.tmpl
var grafanaPrometheusTemplate string

//go:embed templates/observability/grafana/datasource.yml.tmpl
var grafanaDatasourceTemplate string

//go:embed templates/Dockerfile.tmpl
var dockerfileTemplate string

var files = map[string]string{}

func (m AppGenerator) Generate(ebrickConfig AppConfig) {

	files = make(map[string]string)
	files["application.yaml"] = applicationTemplate
	files["cmd/main.go"] = mainTemplate
	files["docker-compose.yml"] = dockerComposeTemplate
	files["go.mod"] = goModTemplate
	files["README.md"] = readmeTemplate
	files["Dockerfile"] = dockerfileTemplate

	if ebrickConfig.Observability {
		files["observability/prometheus/prometheus.yml"] = grafanaPrometheusTemplate
		files["observability/grafana/datasource.yml"] = grafanaDatasourceTemplate
	}
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
	for file, template := range files {
		utils.GenerateFileFromTemplate(file, m.appConfig, template)
		fmt.Println("Generated", file, "successfully.")
	}
}

func (m AppGenerator) postGenerated() {
	fmt.Println("Running post generation tasks...")

	// Run go mod tidy
	utils.ExecCommand("go", "mod", "tidy")

}
