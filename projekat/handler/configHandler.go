package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projekat/model"
	"projekat/service"

	"github.com/gorilla/mux"
)

type ConfigHandler struct {
	service service.ConfigurationService
}

func NewConfigHandler(service service.ConfigurationService) ConfigHandler {
	return ConfigHandler{
		service: service,
	}
}

// GET /configs/{name}/{version}
// GET /configs/{name}/{version}
func (c ConfigHandler) Get(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(10 * time.Second)
	name := mux.Vars(r)["name"]
	version := mux.Vars(r)["version"]

	config, err := c.service.GetConfig(name, version)
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
func (c ConfigHandler) Add(w http.ResponseWriter, r *http.Request) {

	var newConfig model.Configuration
	err := json.NewDecoder(r.Body).Decode(&newConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.AddConfig(newConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Configuration successfully added")
}

func (c ConfigHandler) Delete(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]
	version := mux.Vars(r)["version"]

	err := c.service.DeleteConfig(name, version)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Configuration successfully deleted")

}
