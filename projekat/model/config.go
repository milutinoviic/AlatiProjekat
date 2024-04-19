package model

type Configuration struct {
	Id      string
	Name    string
	Version string
	Params  map[string]string
}

type ConfigurationRepository interface {
	// todo: dodati metodE
	AddConfig(config Configuration) error
	GetConfig(id string) (Configuration, error)
	DeleteConfig(id string) error
}
