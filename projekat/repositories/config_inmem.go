package repositories

import (
	"fmt"

	"projekat/model"
)

type ConfigInMemRepository struct {
	configs map[string]model.Configuration
}

// todo: dodaj implementaciju metoda iz interfejsa ConfigRepository

func NewConfigInMemRepository() model.ConfigurationRepository {
	return &ConfigInMemRepository{
		configs: make(map[string]model.Configuration),
	}
}
func (repo *ConfigInMemRepository) AddConfig(config model.Configuration) error {
	key := config.Name + ":" + config.Version
	if _, ok := repo.configs[key]; ok {
		return fmt.Errorf("Configuration with the same name and version already exists")
	}

	repo.configs[key] = config
	return nil
}

func (repo *ConfigInMemRepository) GetConfig(name string, version string) (model.Configuration, error) {

	key := name + ":" + version
	config, ok := repo.configs[key]
	if !ok {
		return model.Configuration{}, fmt.Errorf("Configuration not found")
	}
	return config, nil

}
func (repo *ConfigInMemRepository) DeleteConfig(name string, version string) error {
	key := name + ":" + version
	if _, ok := repo.configs[key]; ok {
		delete(repo.configs, key)
		return nil
	}
	return fmt.Errorf("Configuration not found")
}
