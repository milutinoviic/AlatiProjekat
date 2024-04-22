package main

import (
	"context"
	"log"

	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"
	"projekat/handler"

	"projekat/repositories"

	"projekat/service"

	"github.com/gorilla/mux"
)

func main() {

	repo := repositories.NewConfigInMemRepository()
	service1 := service.NewConfigService(repo)

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
	router.HandleFunc("/configgroups/{grupName}/{grupVersion}/add", handlerGrup.AddConfigToGroup).Methods("POST")
	router.HandleFunc("/configgroups/{name}/{version}/{configName}/{configVersion}", handlerGrup.RemoveConfigFromGroup).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v\n", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}

	log.Println("Server successfully shut down.")
}
