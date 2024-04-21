package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projekat/model"
	"projekat/service"

	"github.com/gorilla/mux"
)

type ConfigGroupHandler struct {
	service service.ConfigurationGrupService
}

func NewConfigGroupHandler(service service.ConfigurationGrupService) ConfigGroupHandler {
	return ConfigGroupHandler{
		service: service,
	}
}

func (c ConfigGroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	version := mux.Vars(r)["version"]

	config, err := c.service.GetConfigGrup(name, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	resp, err := json.Marshal(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (c ConfigGroupHandler) AddGroup(w http.ResponseWriter, r *http.Request) {

	var newConfigGroup model.ConfigurationGroup
	err := json.NewDecoder(r.Body).Decode(&newConfigGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.AddConfigGrup(newConfigGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Configuration successfully added")
}

func (c ConfigGroupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]
	version := mux.Vars(r)["version"]

	err := c.service.DeleteConfigGrup(name, version)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Configuration successfully deleted")

}
