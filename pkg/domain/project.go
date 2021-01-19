package domain

import (
	"net/http"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	OwnerID     int    `json:"OwnerID"               gorm:"not null"`
	IsDefault   bool   `json:"isDefault"             gorm:"default:false"`
	Title       string `json:"title"                 gorm:"not null"`
	Description string `json:"description,omitempty" db:"description"`
	Logo        string `json:"logo,omitempty"        db:"logo"`
}

type ProjectService interface {
	// CreateProject
	CreateProject(w http.ResponseWriter, r *http.Request)

	// UpdateProject
	UpdateProject(w http.ResponseWriter, r *http.Request)

	// DeleteProject
	DeleteProject(w http.ResponseWriter, r *http.Request)
}

type ProjectRepository interface {
	// InsertProject creates a new project record in the DB
	InsertProject(project Project) (Project, error)

	// UpdateProject updates an existing project record in the DB
	UpdateProject(project Project) (Project, error)

	// DeleteProject deletes an existing project record in the DB
	DeleteProject(id int) error
}
