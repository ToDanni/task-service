package domain

import (
	"net/http"

	"gorm.io/gorm"
)

type Label struct {
	Title       string `json:"title"                  gorm:"not null"`
	Description string `json:"description,omitempty"  db:"description"`
	Colour      string `json:"colour"                 db:"colour"`
	OwnerID     int    `json:"creator"                db:"creator"`
	gorm.Model
}

type LabelService interface {
	// CreateLabel
	Create(w http.ResponseWriter, r *http.Request)

	// UpdateLabel
	Update(w http.ResponseWriter, r *http.Request)

	// DeleteLabel
	Delete(w http.ResponseWriter, r *http.Request)

	// ListLabelsByProject

}

type LabelRepository interface {
	// InsertLabel creates a new label record in the DB
	Insert(label Label) (Label, error)

	// UpdateLabel updates and existing label record in the DB
	Update(label Label) (Label, error)

	// DeleteLabel deletes an existing label record in the DB
	Delete(id int) error
}
