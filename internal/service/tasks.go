package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/task-service/pkg/task"
)

func (s *service) CreateTask(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var taskRequest task.Task
	err = json.Unmarshal(reqBody, &taskRequest)
	if err != nil {
		log.Error(err)
	}

	// Creator ID from the request is overridden
	// so no one can create tasks in place of another person
	userID := r.Context().Value("user_id")
	taskRequest.Creator = userID.(int)

	var createdTask task.Task
	createdTask, err = s.repo.InsertTask(taskRequest)
	marshalled, err := json.Marshal(createdTask)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = writeSuccess(w, marshalled)
	if err != nil {
		log.Error(err)
	}
}

func (s *service) UpdateTask(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	reqID := pathParams["id"]
	id, err := strconv.Atoi(reqID)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		return
	}

	var t task.Task
	err = json.Unmarshal(reqBody, &t)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	t.ID = uint(id)

	updatedTask, err := s.repo.UpdateTask(t)
	marshalled, err := json.Marshal(updatedTask)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = writeSuccess(w, marshalled)
	if err != nil {
		log.Error(err)
	}
}

func (s *service) DeleteTask(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	reqID := pathParams["id"]
	id, err := strconv.Atoi(reqID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	err = s.repo.DeleteTask(id)
	if err != nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	err = writeSuccess(w, []byte(""))
	if err != nil {
		log.Error(err)
	}
}
