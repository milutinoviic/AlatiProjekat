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
	service := service.NewConfigService(repo)
	params := make(map[string]string)
	params["username"] = "pera"
	params["port"] = "5432"
	config := model.Configuration{
		Name:    "db_config",
		Version: "2",
		Params:  params,
	}
	service.AddConfig(config)
	handler := handler.NewConfigHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/configs/{name}/{version}", handler.Get).Methods("GET")
	router.HandleFunc("/configs", handler.Add).Methods("POST")
	router.HandleFunc("/configs/{name}/{version}", handler.Delete).Methods("DELETE")

	http.ListenAndServe("0.0.0.0:8000", router)

}
