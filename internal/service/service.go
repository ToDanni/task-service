package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	repo "github.com/todanni/task-service/internal/repository"
	"github.com/todanni/task-service/pkg/task"
)

type service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) List(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	projID := pathParams["project"]
	projectID, err := strconv.Atoi(projID)

	var tsks []task.Task
	tsks, err = s.repo.SelectTasksByProjectID(projectID)
	if err != nil {
		http.Error(w, "No tasks found", http.StatusNotFound)
		return
	}

	// Marshal response
	marshalled, err := json.Marshal(tsks)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = writeSuccess(w, marshalled)
	if err != nil {
		log.Error(err)
	}
}

func writeSuccess(w http.ResponseWriter, r []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(r)
	return err
}
