package model

type Configuration struct {
	Name    string
	Version string
	Params  map[string]string
}

type ConfigurationRepository interface {
	// todo: dodati metodE
	AddConfig(config Configuration) error
	GetConfig(name string, version string) (Configuration, error)
	DeleteConfig(name string, value string) error
}
