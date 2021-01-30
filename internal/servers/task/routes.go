package task

import "net/http"

func (s *taskService) routes() {
	taskAPI := s.router.PathPrefix("/api/task").Subrouter()

	// POST handles the creation of tasks endpoint
	taskAPI.HandleFunc("/", s.Create).Methods(http.MethodPost)

	// GET handles the listing of tasks endpoint
	taskAPI.HandleFunc("/", s.GetAll).Queries("project", "{project}").Methods(http.MethodGet)

	// PATCH handles the update task endpoint
	taskAPI.HandleFunc("/{id}", s.Update).Methods(http.MethodPatch)

	// DELETE handles the delete task endpoint
	taskAPI.HandleFunc("/{id}", s.Delete).Methods(http.MethodDelete)
}
