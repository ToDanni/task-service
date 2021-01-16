package repository

import (
	"github.com/todanni/task-service/pkg/task"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertTask(task task.Task) (task.Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *repository) SelectTasksByProjectID(projectID int) (tasks []task.Task, err error) {
	err = r.db.Where("project=?", projectID).Find(&tasks).Error
	return tasks, err
}

func (r *repository) UpdateTask(task task.Task) (task.Task, error) {
	// Note: for some stupid reason, update, unlike the rest of the calls from the GORM
	// doesn't update the task after the query with the data from the DB.
	// So I instead perform a query to get the updated object.
	err := r.db.Model(&task).Updates(&task).Error
	r.db.First(&task, task.ID)
	return task, err
}

func (r *repository) DeleteTask(id int) error {
	err := r.db.Delete(&task.Task{}, id).Error
	return err
}
