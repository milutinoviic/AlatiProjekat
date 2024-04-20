package main

import (
	"fmt"
	"projekat/model"

	"projekat/repositories"

	"projekat/service"
)

func main() {
	config := model.Configuration{
		Name:    "Config1",
		Version: "v1",
		Params: map[string]string{
			"param1": "value1",
			"param2": "value2",
		},
	}
	config1 := model.Configuration{
		Name:    "Config223",
		Version: "v2",
		Params: map[string]string{
			"param1": "value1",
			"param2": "value2",
		},
	}

	config2 := model.Configuration{
		Name:    "Config311",
		Version: "v2",
		Params: map[string]string{
			"param1": "value2",
			"param":  "value2",
		},
	}
	config3 := model.Configuration{
		Name:    "Config3cfed",
		Version: "v2",
		Params: map[string]string{
			"param1": "value2",
			"param":  "value2",
		},
	}

	config4 := model.Configuration{
		Name:    "Config3hgdhfg",
		Version: "v2",
		Params: map[string]string{
			"param1": "value2",
			"param":  "value2",
		},
	}

	labelMap := make(map[string]string)
	labelMap["label1"] = "value1"
	labelMap["label2"] = "value2"
	configGroup1 := model.ConfigurationGroup{

		Name:    "Configuration Group 1",
		Version: "v1",
		Configs: []model.Configuration{config},
		Labels:  labelMap,
	}

	configGroup2 := model.ConfigurationGroup{

		Name:    "Sara Group 1",
		Version: "v1",
		Configs: []model.Configuration{config},
		Labels:  labelMap,
	}

	configGroup3 := model.ConfigurationGroup{

		Name:    "Stefan Group 1",
		Version: "v1",
		Configs: []model.Configuration{config},
		Labels:  labelMap,
	}
	configGroup4 := model.ConfigurationGroup{

		Name:    "Tijana Group 1",
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
	service1.AddConfig(config2)
	service1.AddConfig(config3)
	service1.AddConfig(config4)
	service1.DeleteConfig(config1.Name, config1.Version)
	service1.DeleteConfig(config3.Name, config3.Version)
	//service1.DeleteConfig(config4.Name, config4.Version)

	fmt.Println(service1.GetConfig(config1.Name, config1.Version))
	fmt.Println(service1.GetConfig(config2.Name, config2.Version))
	fmt.Println(service1.GetConfig(config3.Name, config3.Version))
	fmt.Println(service1.GetConfig(config4.Name, config4.Version))
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("---------------------------------------------------------------")
	serviceGrup.AddConfigGrup(configGroup1)
	serviceGrup.AddConfigGrup(configGroup2)
	serviceGrup.AddConfigGrup(configGroup3)
	serviceGrup.AddConfigGrup(configGroup4)
	serviceGrup.DeleteConfigGrup(configGroup1.Name, configGroup1.Version)
	serviceGrup.DeleteConfigGrup(configGroup3.Name, configGroup3.Version)
	fmt.Println(serviceGrup.GetConfigGrup(configGroup1.Name, configGroup1.Version))
	fmt.Println(serviceGrup.GetConfigGrup(configGroup2.Name, configGroup2.Version))
	fmt.Println(serviceGrup.GetConfigGrup(configGroup3.Name, configGroup3.Version))
	fmt.Println(serviceGrup.GetConfigGrup(configGroup4.Name, configGroup4.Version))
	//fmt.Println(serviceGrup.GetConfigGrup(configGroup1.Name, config.Version))

}
