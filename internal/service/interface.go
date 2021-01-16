package service

import (
	"net/http"
)

type Service interface {
	// ListAll
	List(w http.ResponseWriter, r *http.Request)

	// Task interface endpoints
	CreateTask(w http.ResponseWriter, r *http.Request)

	UpdateTask(w http.ResponseWriter, r *http.Request)

	DeleteTask(w http.ResponseWriter, r *http.Request)

	// Project interface endpoints
	CreateProject(w http.ResponseWriter, r *http.Request)

	UpdateProject(w http.ResponseWriter, r *http.Request)

	DeleteProject(w http.ResponseWriter, r *http.Request)

	// Label interface endpoints
	CreateLabel(w http.ResponseWriter, r *http.Request)

	UpdateLabel(w http.ResponseWriter, r *http.Request)

	DeleteLabel(w http.ResponseWriter, r *http.Request)
}
