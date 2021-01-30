package task

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/task-service/pkg/domain"
)

func NewTaskService(repo domain.TaskRepository, router mux.Router) domain.TaskService {
	server := &taskService{
		repo:   repo,
		router: &router,
	}
	server.routes()
	return server
}

type taskService struct {
	repo   domain.TaskRepository
	router *mux.Router
}

func (s *taskService) Create(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var taskRequest domain.Task
	err = json.Unmarshal(reqBody, &taskRequest)
	if err != nil {
		log.Error(err)
	}

	// Creator ID from the request is overridden
	// so no one can create tasks in place of another person
	userID := r.Context().Value("user_id")
	taskRequest.OwnerID = userID.(int)

	var createdTask domain.Task

	createdTask, err = s.repo.Insert(taskRequest)
	marshalled, err := json.Marshal(createdTask)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(marshalled)
	if err != nil {
		log.Error(err)
	}
}

func (s *taskService) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *taskService) GetAll(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *taskService) Update(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	reqID := pathParams["id"]
	id, err := strconv.Atoi(reqID)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		return
	}

	var t domain.Task
	err = json.Unmarshal(reqBody, &t)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	t.ID = uint(id)

	updatedTask, err := s.repo.Update(t)
	marshalled, err := json.Marshal(updatedTask)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(marshalled)
	if err != nil {
		log.Error(err)
	}
}

func (s *taskService) Delete(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	reqID := pathParams["id"]
	id, err := strconv.Atoi(reqID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	err = s.repo.Delete(id)
	if err != nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(""))
	if err != nil {
		log.Error(err)
	}
}
