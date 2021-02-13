package label

import "net/http"

func (s *labelService) routes() {
	// Define the prefix for the labels API
	labelAPI := s.router.PathPrefix("/api/label").Subrouter()

	// GET lists all the labels for a project
	labelAPI.HandleFunc("/", s.List).Queries("project", "{project}").Methods(http.MethodGet)

	// POST handler for the Create labels endpoint
	labelAPI.HandleFunc("/", s.Create).Methods(http.MethodPost)

	// PATCH handler for the Update label endpoint
	labelAPI.HandleFunc("/{id}", s.Update).Methods(http.MethodPatch)

	// DELETE handler for the Delete labels endpoint
	labelAPI.HandleFunc("/{id}", s.Delete).Methods(http.MethodDelete)
}
