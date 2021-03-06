package domain

import (
	"net/http"

	"gorm.io/gorm"
)

type Label struct {
	ProjectID uint `json:"ProjectID"`
	gorm.Model
	EditableFields
}

type EditableFields struct {
	Title       string `json:"Title"                  gorm:"not null"`
	Description string `json:"Description,omitempty"  db:"description"`
	Colour      string `json:"Colour"                 db:"colour"`
}

type LabelService interface {
	// CreateLabel
	Create(w http.ResponseWriter, r *http.Request)

	// UpdateLabel
	Update(w http.ResponseWriter, r *http.Request)

	// DeleteLabel
	Delete(w http.ResponseWriter, r *http.Request)

	// ListLabelsByProject
	List(w http.ResponseWriter, r *http.Request)
}

type LabelRepository interface {
	// Insert creates a new label record in the DB
	Insert(label Label) (Label, error)

	// Update updates and existing label record in the DB
	Update(label Label) (Label, error)

	// Delete deletes an existing label record in the DB
	Delete(id int) error

	// SelectByProject returns all records in the DB of labels for a given project
	SelectByProject(projectID int) ([]Label, error)
}
