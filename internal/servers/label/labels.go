package label

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/task-service/pkg/domain"
)

type labelService struct {
	repo   domain.LabelRepository
	router *mux.Router
}

func NewLabelService(repo domain.LabelRepository, router mux.Router) domain.LabelService {
	server := &labelService{
		repo:   repo,
		router: &router,
	}
	server.routes()
	return server
}

func (s *labelService) Create(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var labelRequest domain.Label
	err = json.Unmarshal(reqBody, &labelRequest)
	if err != nil {
		log.Error(err)
	}

	// Creator ID from the request is overridden
	// so no one can create tasks in place of another person
	userID := r.Context().Value("user_id")
	labelRequest.OwnerID = userID.(int)

	var createdLabel domain.Label
	createdLabel, err = s.repo.Insert(labelRequest)
	marshalled, err := json.Marshal(createdLabel)
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

func (s *labelService) Update(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	reqID := pathParams["id"]
	id, err := strconv.Atoi(reqID)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		return
	}

	var l domain.Label
	err = json.Unmarshal(reqBody, &l)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	l.ID = uint(id)

	updatedLabel, err := s.repo.Update(l)
	marshalled, err := json.Marshal(updatedLabel)
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

func (s *labelService) Delete(w http.ResponseWriter, r *http.Request) {
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

func (s *labelService) List(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	reqID := pathParams["project"]
	projectID, err := strconv.Atoi(reqID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	var labels []domain.Label
	labels, err = s.repo.SelectByProject(projectID)
	if err != nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	marshalled, err := json.Marshal(labels)
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
