package repository

import "github.com/todanni/task-service/pkg/task"

type Repository interface {
	// InsertTask persists a task in the DB
	InsertTask(task task.Task) (task.Task, error)

	// SelectTasksByProjectID returns all records in the DB of tasks filtered by project
	SelectTasksByProjectID(projectID int) (tasks []task.Task, err error)

	// UpdateTask updates a given task in the DB
	UpdateTask(task task.Task) (task.Task, error)

	// DeleteTask deletes a task record from the DB
	DeleteTask(id int) error
}
