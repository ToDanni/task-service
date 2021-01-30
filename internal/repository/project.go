package repository

import (
	"github.com/todanni/task-service/pkg/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) SelectProjectsByUser(userID int) ([]domain.Project, error) {
	var userProjects []domain.Project
	err := r.db.Find(&userProjects, "user_id = ? ", userID).Error
	return userProjects, err
}

func (r *repository) InsertProject(project domain.Project) (domain.Project, error) {
	err := r.db.Create(&project).Error
	return project, err
}

func (r *repository) UpdateProject(project domain.Project) (domain.Project, error) {
	err := r.db.Model(&project).Updates(&project).Error
	r.db.First(&project, project.ID)
	return project, err
}

func (r *repository) DeleteProject(id int) error {
	err := r.db.Delete(&domain.Project{}, id).Error
	return err
}

func NewProjectRepository(db *gorm.DB) domain.ProjectRepository {
	return &repository{
		db: db,
	}
}
