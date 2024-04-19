package repositories

import (
	"errors"

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

func (repo *ConfigGrupInMemRepository) GetConfigGrupe(id string) (model.ConfigurationGroup, error) {
	grupConfig, ok := repo.groups[id]

	if !ok {
		return model.ConfigurationGroup{}, errors.New("Not found ConfigGrup")

	}

	return grupConfig, nil
}

func (repo *ConfigGrupInMemRepository) AddConfigGrup(groupConfig model.ConfigurationGroup) error {
	repo.groups[groupConfig.Id] = groupConfig
	return nil
}

func (repo *ConfigGrupInMemRepository) DeleteConfigGrup(id string) error {
	_, exist := repo.groups[id]
	if !exist {
		return errors.New("Not found group")
	}

	delete(repo.groups, id)

	return nil

}
