package service

import (
	"fmt"

	"projekat/model"
)

type ConfigurationService struct {
	repo model.ConfigurationRepository
}

func NewConfigService(repo model.ConfigurationRepository) ConfigurationService {
	var service ConfigurationService

	service.repo = repo

	return service

}

func (s ConfigurationService) Hello() {
	fmt.Println("hello from config service")
}

func (s ConfigurationService) GetConfig(id string) (model.Configuration, error) {
	return s.repo.GetConfig(id)

}

func (s ConfigurationService) AddConfig(config model.Configuration) error {
	return s.repo.AddConfig(config)

}
func (s ConfigurationService) DeleteConfig(id string) error {
	return s.repo.DeleteConfig(id)
}
