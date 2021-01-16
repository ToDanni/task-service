package task

import (
	"github.com/todanni/task-service/pkg/label"
	"github.com/todanni/task-service/pkg/project"
	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Labels      []label.Label `gorm:"many2many:task_labels;"`
	Project     project.Project
	ProjectID   int
	OwnerID     int         `json:"creator"     db:"creator"`
	Assignee    int         `json:"assignee"    db:"assignee"`
	Done        bool        `json:"done"        gorm:"default:false"`
	Title       string      `json:"title"       gorm:"not null"`
	Status      string      `json:"status"      gorm:"default:Todo"`
	CompletedAt null.String `json:"completedAt" db:"completed_at"`
	Deadline    null.String `json:"deadline"    db:"deadline"`
	Description null.String `json:"description" db:"description"`
}
