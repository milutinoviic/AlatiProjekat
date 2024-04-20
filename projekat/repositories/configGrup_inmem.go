package repositories

import (
	"fmt"
	"projekat/model"
)

type ConfigGrupInMemRepository struct {
	groups map[string]model.ConfigurationGroup
}

func NewConfigGrupInMemRepository() model.ConfigurationGroupRepository {
	return &ConfigGrupInMemRepository{

		groups: make(map[string]model.ConfigurationGroup),
	}
}

func (repo *ConfigGrupInMemRepository) GetConfigGrupe(name string, version string) (model.ConfigurationGroup, error) {
	key := name + ":" + version
	configGrup, ok := repo.groups[key]
	if !ok {
		return model.ConfigurationGroup{}, fmt.Errorf("Configuration Grup not found")
	}
	return configGrup, nil

}

func (repo *ConfigGrupInMemRepository) AddConfigGrup(groupConfig model.ConfigurationGroup) error {
	key := groupConfig.Name + ":" + groupConfig.Version
	repo.groups[key] = groupConfig
	return nil

}

func (repo *ConfigGrupInMemRepository) DeleteConfigGrup(name string, version string) error {
	key := name + ":" + version
	if _, ok := repo.groups[key]; ok {
		delete(repo.groups, key)
		return nil
	}
	return fmt.Errorf("Configuration not found")

}

func (repo *ConfigGrupInMemRepository) AddConfigToGroup(groupName string, version string, config model.Configuration) error {
	key := groupName + ":" + version
	group, ok := repo.groups[key]
	if !ok {
		return fmt.Errorf("Configuration Group not found")
	}
	group.Configs = append(group.Configs, config)
	repo.groups[key] = group
	return nil
}
func (repo *ConfigGrupInMemRepository) RemoveConfigFromGroup(groupName string, version string, configName string, configVersion string) error {
	key := groupName + ":" + version
	group, ok := repo.groups[key]
	if !ok {
		return fmt.Errorf("Configuration Group not found")
	}
	var updatedConfigs []model.Configuration
	for _, config := range group.Configs {
		if config.Name != configName && config.Version != configVersion {
			updatedConfigs = append(updatedConfigs, config)
		}
	}
	group.Configs = updatedConfigs
	repo.groups[key] = group
	return nil
}
