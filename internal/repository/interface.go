package repository

import (
	"github.com/todanni/task-service/pkg/label"
	"github.com/todanni/task-service/pkg/project"
	"github.com/todanni/task-service/pkg/task"
)

type Repository interface {
	// InsertTask persists a task in the DB
	InsertTask(task task.Task) (task.Task, error)

	// SelectTasksByProjectID returns all records in the DB of tasks filtered by project
	SelectTasksByProjectID(projectID int) (tasks []task.Task, err error)

	// UpdateTask updates a given task in the DB
	UpdateTask(task task.Task) (task.Task, error)

	// DeleteTask deletes a task record from the DB
	DeleteTask(id int) error

	// SelectAllItemsForUser returns all projects, tasks and labels for a user
	SelectAllItemsForUser() (projects []project.Project, err error)

	// InsertProject creates a new project record in the DB
	InsertProject(project project.Project) (project.Project, error)

	// UpdateProject updates an existing project record in the DB
	UpdateProject(project project.Project) (project.Project, error)

	// DeleteProject deletes an existing project record in the DB
	DeleteProject(id int) error

	// InsertLabel creates a new label record in the DB
	InsertLabel(label label.Label) (label.Label, error)

	// UpdateLabel updates and existing label record in the DB
	UpdateLabel(label label.Label) (label.Label, error)

	// DeleteLabel deletes an existing label record in the DB
	DeleteLabel(id int) error
}
