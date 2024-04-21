package main

import (
	"net/http"
	"projekat/handler"
	"projekat/model"

	"projekat/repositories"

	"projekat/service"

	"github.com/gorilla/mux"
)

func main() {

	repo := repositories.NewConfigInMemRepository()
	service1 := service.NewConfigService(repo)

	params := make(map[string]string)
	params["username"] = "pera"
	params["port"] = "5432"
	config := model.Configuration{
		Name:    "db_config",
		Version: "2",
		Params:  params,
	}

	service1.AddConfig(config)

	handler1 := handler.NewConfigHandler(service1)

	repoGrup := repositories.NewConfigGrupInMemRepository()

	serviceGrup := service.NewConfigGrupService(repoGrup)

	handlerGrup := handler.NewConfigGroupHandler(serviceGrup)

	router := mux.NewRouter()

	router.HandleFunc("/configs/{name}/{version}", handler1.Get).Methods("GET")
	router.HandleFunc("/configs", handler1.Add).Methods("POST")
	router.HandleFunc("/configs/{name}/{version}", handler1.Delete).Methods("DELETE")

	router.HandleFunc("/configgroups/{name}/{version}", handlerGrup.GetGroup).Methods("GET")
	router.HandleFunc("/configgroups", handlerGrup.AddGroup).Methods("POST")
	router.HandleFunc("/configgroups/{name}/{version}", handlerGrup.DeleteGroup).Methods("DELETE")

	http.ListenAndServe("0.0.0.0:8000", router)
}
