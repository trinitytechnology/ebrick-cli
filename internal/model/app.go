package model

type EBrickApp struct {
	Name            string   `yaml:"name"`
	Package         string   `yaml:"package"`
	Database        bool     `yaml:"database"`
	Observability   bool     `yaml:"observability"`
	Cache           bool     `yaml:"cache"`
	Messaging       bool     `yaml:"messaging"`
	Version         string   `yaml:"version"`
	ExternalModules []string `yaml:"extModules"`
	InternalModules []Module `yaml:"intModules"`
}
