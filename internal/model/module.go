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
