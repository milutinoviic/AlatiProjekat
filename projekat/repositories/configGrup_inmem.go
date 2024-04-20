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
