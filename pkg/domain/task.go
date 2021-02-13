package domain

import (
	"net/http"

	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Labels      []Label     `json:"Labels" gorm:"many2many:task_labels;"`
	Project     Project     `json:"-"`
	ProjectID   int         `json:"ProjectID"`
	OwnerID     int         `json:"OwnerID"     db:"creator"`
	AssigneeID  int         `json:"AssigneeID"    db:"assignee"`
	Done        bool        `json:"Done"        gorm:"default:false"`
	Title       string      `json:"Title"       gorm:"not null"`
	Status      string      `json:"Status"      gorm:"default:Todo"`
	CompletedAt null.String `json:"CompletedAt" db:"completed_at"`
	Deadline    null.String `json:"Deadline"    db:"deadline"`
	Description null.String `json:"Description" db:"description"`
}

type TaskService interface {
	// Create
	Create(w http.ResponseWriter, r *http.Request)

	// Get
	Get(w http.ResponseWriter, r *http.Request)

	// GetAll
	GetAll(w http.ResponseWriter, r *http.Request)

	// Update
	Update(w http.ResponseWriter, r *http.Request)

	// Delete
	Delete(w http.ResponseWriter, r *http.Request)
}

type TaskRepository interface {
	// Select returns a DB record for a task by ID
	Select(id int) (Task, error)

	// SelectAll
	SelectAll() ([]Task, error)

	// Insert persists a task in the DB
	Insert(task Task) (Task, error)

	// Update updates a given task in the DB
	Update(task Task) (Task, error)

	// Delete deletes a task record from the DB
	Delete(id int) error
}
