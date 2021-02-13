package domain

import (
	"net/http"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	OwnerID     int     `json:"OwnerID"               gorm:"not null"`
	IsDefault   bool    `json:"IsDefault"             gorm:"default:false"`
	Title       string  `json:"Title"                 gorm:"not null"`
	Description string  `json:"Description,omitempty" db:"description"`
	Logo        string  `json:"Logo,omitempty"        db:"logo"`
	Labels      []Label `json:"Labels"`
	//Users       []int   `json:"Members"`
}

type ProjectService interface {
	// CreateProject
	CreateProject(w http.ResponseWriter, r *http.Request)

	// UpdateProject
	UpdateProject(w http.ResponseWriter, r *http.Request)

	// DeleteProject
	DeleteProject(w http.ResponseWriter, r *http.Request)

	// ListProjects
	ListProjects(w http.ResponseWriter, r *http.Request)
}

type ProjectRepository interface {
	// InsertProject creates a new project record in the DB
	InsertProject(project Project) (Project, error)

	// UpdateProject updates an existing project record in the DB
	UpdateProject(project Project) (Project, error)

	// DeleteProject deletes an existing project record in the DB
	DeleteProject(id int) error

	// SelectProjectsByUser
	SelectProjectsByUser(userID int) ([]Project, error)
}
