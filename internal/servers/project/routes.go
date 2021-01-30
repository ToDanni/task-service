package project

import (
	"net/http"
)

func (s *projectService) routes() {
	projectAPI := s.router.PathPrefix("/api/project").Subrouter()

	// POST
	projectAPI.HandleFunc("/", s.CreateProject).Methods(http.MethodPost)

	// PATCH
	projectAPI.HandleFunc("/{id}", s.UpdateProject).Methods(http.MethodPatch)

	// DELETE
	projectAPI.HandleFunc("/{id}", s.DeleteProject).Methods(http.MethodDelete)
}
