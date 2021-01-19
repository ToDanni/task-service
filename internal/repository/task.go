package repository

import (
	"github.com/todanni/task-service/pkg/domain"
	"gorm.io/gorm"
)

func (r *taskRepository) Select(id int) (domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id).Error
	return task, err
}

// TODO: determine how to query all tasks for a user
func (r *taskRepository) SelectAll() ([]domain.Task, error) {
	panic("implement me")
}

func (r *taskRepository) Insert(task domain.Task) (domain.Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) Update(task domain.Task) (domain.Task, error) {
	// Note: for some stupid reason, update, unlike the rest of the calls from the GORM
	// doesn't update the task after the query with the data from the DB.
	// So I instead perform a query to get the updated object.
	err := r.db.Model(&task).Updates(&task).Error
	r.db.Joins("Project").First(&task, task.ID)
	return task, err
}

func (r *taskRepository) Delete(id int) error {
	err := r.db.Delete(&domain.Task{}, id).Error
	return err
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

type taskRepository struct {
	db *gorm.DB
}
