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

	if _, ok := repo.groups[key]; ok {
		return fmt.Errorf("Configuration group with the same name and version already exists")
	}
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

	for _, existingConfig := range group.Configs {
		if existingConfig.Name == config.Name && existingConfig.Version == config.Version {
			return fmt.Errorf("Configuration with the same name and values already exists in the group")
		}
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

	var found bool
	var updatedConfigs []model.Configuration
	for _, config := range group.Configs {
		if config.Name == configName && config.Version == configVersion {
			found = true
		} else {
			updatedConfigs = append(updatedConfigs, config)
		}
	}

	if !found {
		return fmt.Errorf("Configuration not found in the group")
	}

	group.Configs = updatedConfigs
	repo.groups[key] = group
	return nil
}
