package task

import (
	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Project     int         `json:"project"     db:"project"`
	Creator     int         `json:"creator"     db:"creator"`
	Assignee    int         `json:"assignee"    db:"assignee"`
	Done        bool        `json:"done"        db:"done"`
	Title       string      `json:"title"       db:"title"`
	Status      string      `json:"status"      db:"status"`
	CompletedAt null.String `json:"completedAt" db:"completed_at"`
	Deadline    null.String `json:"deadline"    db:"deadline"`
	Description null.String `json:"description" db:"description"`
}
