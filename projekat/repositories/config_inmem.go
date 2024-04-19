package repositories

import (
	"errors"
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
	repo.configs[config.Id] = config
	return nil

}

func (repo *ConfigInMemRepository) GetConfig(id string) (model.Configuration, error) {
	config, exsist := repo.configs[id]

	if !exsist {
		return model.Configuration{}, fmt.Errorf("Konfiguracija sa ID-om nije pronadjena", id)

	}

	return config, nil

}
func (repo *ConfigInMemRepository) DeleteConfig(id string) error {
	_, exists := repo.configs[id]
	if !exists {
		return errors.New("Not found config")
	}

	delete(repo.configs, id)

	return nil
}
