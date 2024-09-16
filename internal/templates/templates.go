package templates

import (
	_ "embed"
)

var (
	//go:embed templates/application.yaml.tmpl
	ApplicationTemplate string

	//go:embed templates/main.go.tmpl
	MainTemplate string

	//go:embed templates/docker-compose.yml.tmpl
	DockerComposeTemplate string

	//go:embed templates/go.mod.tmpl
	GoModTemplate string

	//go:embed templates/README.md.tmpl
	ReadmeTemplate string

	//go:embed templates/observability/prometheus/prometheus.yml.tmpl
	GrafanaPrometheusTemplate string

	//go:embed templates/observability/grafana/datasource.yml.tmpl
	GrafanaDatasourceTemplate string

	//go:embed templates/Dockerfile.tmpl
	DockerfileTemplate string

	//go:embed templates/module.go.tmpl
	ModuleTemplate string
)
