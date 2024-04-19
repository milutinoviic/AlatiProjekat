package main

import (
	"fmt"
	"projekat/model"

	"projekat/repositories"

	"projekat/service"
)

func main() {
	config := model.Configuration{
		Id:      "1",
		Name:    "Config1",
		Version: "v1",
		Params: map[string]string{
			"param1": "value1",
			"param2": "value2",
		},
	}
	config1 := model.Configuration{
		Id:      "2",
		Name:    "Config2",
		Version: "v2",
		Params: map[string]string{
			"param1": "value1",
			"param2": "value2",
		},
	}
	labelMap := make(map[string]string)
	labelMap["label1"] = "value1"
	labelMap["label2"] = "value2"
	configGroup1 := model.ConfigurationGroup{
		Id:      "group1",
		Name:    "Configuration Group 1",
		Version: "v1",
		Configs: []model.Configuration{config},
		Labels:  labelMap,
	}
	repo := repositories.NewConfigInMemRepository()
	service1 := service.NewConfigService(repo)
	repoGrup := repositories.NewConfigGrupInMemRepository()
	serviceGrup := service.NewConfigGrupService(repoGrup)

	service1.Hello()
	service1.AddConfig(config)
	service1.AddConfig(config1)
	fmt.Println(service1.GetConfig(config.Id))

	serviceGrup.AddConfigGrup(configGroup1)
	//fmt.Println(serviceGrup.GetConfigGrup(configGroup1.Id))

}
