package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/task-service/internal/config"
	"github.com/todanni/task-service/internal/database"
	"github.com/todanni/task-service/internal/repository"
	"github.com/todanni/task-service/internal/service"
)

func main() {
	// Read config
	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Setup database
	db, err := database.Open(cfg)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Create service
	var svc service.Service
	svc = service.NewService(repository.NewRepository(db))

	// Setup router
	r := mux.NewRouter()
	api := r.PathPrefix("/api/task").Subrouter()
	api.HandleFunc("/", svc.Create).Methods(http.MethodPost)
	api.HandleFunc("/", svc.List).Queries("project", "{project}").Methods(http.MethodGet)
	api.HandleFunc("/{id}", svc.Update).Methods(http.MethodPatch)
	api.HandleFunc("/{id}", svc.Delete).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8083", r))

}
