package main

import (
	"net/http"
	"os"

	"github.com/todanni/task-service/pkg/domain"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/task-service/internal/config"
	"github.com/todanni/task-service/internal/database"
	"github.com/todanni/task-service/internal/middleware"
	"github.com/todanni/task-service/internal/repository"
	"github.com/todanni/task-service/internal/servers/label"
	"github.com/todanni/task-service/internal/servers/project"
	"github.com/todanni/task-service/internal/servers/task"
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

	err = db.AutoMigrate(&domain.Project{}, &domain.Label{})
	if err != nil {
		log.Error(err)
	}

	// Setup router
	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	r.Use(middleware.LoggingMiddleware)

	// Create task service
	task.NewTaskService(repository.NewTaskRepository(db), r)

	// Create project service
	project.NewProjectService(repository.NewProjectRepository(db), r)

	// Create label service
	label.NewLabelService(repository.NewLabelRepository(db), r)

	log.Fatal(http.ListenAndServe(":8083", r))
}
