package model

type Module struct {
	Id          string `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Package     string `yaml:"package"`
	External    bool   `yaml:"external"`
	Rest        bool   `yaml:"rest"`
	Graphql     bool   `yaml:"graphql"`
	Grpc        bool   `yaml:"grpc"`
	Messaging   bool   `yaml:"messaging"`
	Auth        bool   `yaml:"auth"`
}
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
