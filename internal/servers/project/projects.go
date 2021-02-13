package project

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/task-service/pkg/domain"
)

type projectService struct {
	repo   domain.ProjectRepository
	router *mux.Router
}

func NewProjectService(repo domain.ProjectRepository, router *mux.Router) domain.ProjectService {
	server := &projectService{
		repo:   repo,
		router: router,
	}
	// Setup routing for the server
	server.routes()
	return server
}

func (s *projectService) CreateProject(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var projectRequest domain.Project
	err = json.Unmarshal(reqBody, &projectRequest)
	if err != nil {
		log.Error(err)
	}

	// Creator ID from the request is overridden
	// so no one can create tasks in place of another person
	userID := r.Context().Value("user_id")
	projectRequest.OwnerID = userID.(int)

	var createdProject domain.Project
	createdProject, err = s.repo.InsertProject(projectRequest)
	marshalled, err := json.Marshal(createdProject)
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

func (s *projectService) UpdateProject(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	reqID := pathParams["id"]
	id, err := strconv.Atoi(reqID)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		return
	}

	var p domain.Project
	err = json.Unmarshal(reqBody, &p)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	p.ID = uint(id)

	updatedProject, err := s.repo.UpdateProject(p)
	marshalled, err := json.Marshal(updatedProject)
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

func (s *projectService) DeleteProject(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	reqID := pathParams["id"]
	id, err := strconv.Atoi(reqID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	err = s.repo.DeleteProject(id)
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

func (s *projectService) ListProjects(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")
	// userID.(int)

	var projects []domain.Project
	projects, err := s.repo.SelectProjectsByUser(userID.(int))
	if err != nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	marshalled, err := json.Marshal(projects)
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
