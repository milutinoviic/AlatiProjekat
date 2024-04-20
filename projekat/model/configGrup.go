package model

// import "../mo"

type ConfigurationGroup struct {
	Name    string
	Version string
	Configs []Configuration
	Labels  map[string]string
}

type ConfigurationGroupRepository interface {
	AddConfigGrup(group ConfigurationGroup) error
	GetConfigGrupe(name string, version string) (ConfigurationGroup, error)
	DeleteConfigGrup(name string, version string) error
	AddConfigToGroup(groupName string, version string, config Configuration) error
	RemoveConfigFromGroup(groupName string, version string, configName string, configVersion string) error
}
