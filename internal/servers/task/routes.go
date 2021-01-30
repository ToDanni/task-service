package task

import "net/http"

func (s *taskService) routes() {
	// POST handles the creation of tasks endpoint
	s.router.HandleFunc("/api/task/", s.Create).Methods(http.MethodPost)

	// GET handles the listing of tasks endpoint
	s.router.HandleFunc("/api/task/", s.GetAll).Queries("project", "{project}").Methods(http.MethodGet)

	// PATCH handles the update task endpoint
	s.router.HandleFunc("/api/task/{id}", s.Update).Methods(http.MethodPatch)

	// DELETE handles the delete task endpoint
	s.router.HandleFunc("/api/task/{id}", s.Delete).Methods(http.MethodDelete)
}
