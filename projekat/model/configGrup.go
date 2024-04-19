package model

// import "../mo"

type ConfigurationGroup struct {
	Id      string
	Name    string
	Version string
	Configs []Configuration
	Labels  map[string]string
}

type ConfigurationGroupRepository interface {
	AddConfigGrup(group ConfigurationGroup) error
	GetConfigGrupe(id string) (ConfigurationGroup, error)
	DeleteConfigGrup(id string) error
}
