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
	"github.com/todanni/task-service/pkg/label"
	"github.com/todanni/task-service/pkg/project"
	"github.com/todanni/task-service/pkg/task"
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

	// Auto-migrate
	err = db.AutoMigrate(&project.Project{}, &task.Task{}, &label.Label{})
	if err != nil {
		log.Error(err)
	}

	proj := project.Project{
		OwnerID:     1,
		IsDefault:   true,
		Title:       "Test Project",
		Description: "Test Description",
		Logo:        "http://imgr.com/borken.png",
	}

	db.Create(&proj)

	t := task.Task{
		ProjectID: 1,
		OwnerID:   1,
		Assignee:  1,
		Done:      false,
		Title:     "Test Task",
		Status:    "In Progress",
		Project:   proj,
	}
	db.Create(&t)

	l := label.Label{
		Title:       "Test Label",
		Description: "Test Description",
		Colour:      "red",
		OwnerID:     1,
	}
	db.Create(&l)

	// Create service
	var svc service.Service
	svc = service.NewService(repository.NewRepository(db))

	// Setup router
	r := mux.NewRouter()
	//r.Use(middleware.Middleware)
	//r.Use(middleware.LoggingMiddleware)

	// Task endpoints
	api := r.PathPrefix("/api/task").Subrouter()
	api.HandleFunc("/", svc.CreateTask).Methods(http.MethodPost)
	api.HandleFunc("/", svc.List).Queries("project", "{project}").Methods(http.MethodGet)
	api.HandleFunc("/{id}", svc.UpdateTask).Methods(http.MethodPatch)
	api.HandleFunc("/{id}", svc.DeleteTask).Methods(http.MethodDelete)

	// Project endpoints
	projectAPI := r.PathPrefix("/api/project").Subrouter()
	projectAPI.HandleFunc("/", svc.CreateProject).Methods(http.MethodPost)
	projectAPI.HandleFunc("/{id}", svc.UpdateProject).Methods(http.MethodPatch)
	projectAPI.HandleFunc("/{id}", svc.DeleteProject).Methods(http.MethodDelete)

	// Label endpoints
	labelAPI := r.PathPrefix("/api/label").Subrouter()
	labelAPI.HandleFunc("/", svc.CreateLabel).Methods(http.MethodPost)
	labelAPI.HandleFunc("/{id}", svc.UpdateLabel).Methods(http.MethodPatch)
	labelAPI.HandleFunc("/{id}", svc.DeleteLabel).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8083", r))

}
