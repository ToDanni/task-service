package main

import (
	"net/http"
	"os"

	"github.com/todanni/task-service/pkg/domain"

	"github.com/todanni/task-service/internal/middleware"

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

	// Auto-migrate
	err = db.AutoMigrate(&domain.Project{}, &domain.Task{}, &domain.Label{})
	if err != nil {
		log.Error(err)
	}

	// Setup router
	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	r.Use(middleware.LoggingMiddleware)

	var taskService domain.TaskService
	taskService = service.NewTaskService(repository.NewTaskRepository(db))
	taskAPI := r.PathPrefix("/api/task").Subrouter()
	taskAPI.HandleFunc("/", taskService.Create).Methods(http.MethodPost)
	taskAPI.HandleFunc("/", taskService.GetAll).Queries("project", "{project}").Methods(http.MethodGet)
	taskAPI.HandleFunc("/{id}", taskService.Update).Methods(http.MethodPatch)
	taskAPI.HandleFunc("/{id}", taskService.Delete).Methods(http.MethodDelete)

	var projectService domain.ProjectService
	projectService = service.NewProjectService(repository.NewProjectRepository(db))
	projectAPI := r.PathPrefix("/api/project").Subrouter()
	projectAPI.HandleFunc("/", projectService.CreateProject).Methods(http.MethodPost)
	projectAPI.HandleFunc("/{id}", projectService.UpdateProject).Methods(http.MethodPatch)
	projectAPI.HandleFunc("/{id}", projectService.DeleteProject).Methods(http.MethodDelete)

	var labelService domain.LabelService
	labelService = service.NewLabelService(repository.NewLabelRepository(db))
	labelAPI := r.PathPrefix("/api/label").Subrouter()
	labelAPI.HandleFunc("/", labelService.Create).Methods(http.MethodPost)
	labelAPI.HandleFunc("/{id}", labelService.Update).Methods(http.MethodPatch)
	labelAPI.HandleFunc("/{id}", labelService.Delete).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8083", r))
}
