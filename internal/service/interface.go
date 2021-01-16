package service

import (
	"net/http"
)

type Service interface {
	// List tasks for a given project
	List(w http.ResponseWriter, r *http.Request)

	// Create a task
	Create(w http.ResponseWriter, r *http.Request)

	// Update a task
	Update(w http.ResponseWriter, r *http.Request)

	// Delete a task
	Delete(w http.ResponseWriter, r *http.Request)
}
