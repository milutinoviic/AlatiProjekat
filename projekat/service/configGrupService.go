package service

import (
	"projekat/model"
)

type ConfigurationGrupService struct {
	repo model.ConfigurationGroupRepository
}

func NewConfigGrupService(repo model.ConfigurationGroupRepository) ConfigurationGrupService {
	var service ConfigurationGrupService

	service.repo = repo

	return service

}

func (s ConfigurationGrupService) AddConfigGrup(group model.ConfigurationGroup) error {
	return s.repo.AddConfigGrup(group)
}

func (s ConfigurationGrupService) GetConfigGrup(id string) (model.ConfigurationGroup, error) {
	return s.repo.GetConfigGrupe(id)
}

func (s ConfigurationGrupService) DeleteConfigGrup(id string) error {
	return s.repo.DeleteConfigGrup(id)
}
