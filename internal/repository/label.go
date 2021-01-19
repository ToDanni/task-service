package repository

import (
	"github.com/todanni/task-service/pkg/domain"
	"gorm.io/gorm"
)

func (r *labelRepository) Insert(label domain.Label) (domain.Label, error) {
	err := r.db.Create(&label).Error
	return label, err
}

func (r *labelRepository) Update(label domain.Label) (domain.Label, error) {
	err := r.db.Model(&label).Updates(&label).Error
	r.db.First(&label, label.ID)
	return label, err
}

func (r *labelRepository) Delete(id int) error {
	err := r.db.Delete(&domain.Label{}, id).Error
	return err
}

func NewLabelRepository(db *gorm.DB) domain.LabelRepository {
	return &labelRepository{
		db: db,
	}
}

type labelRepository struct {
	db *gorm.DB
}
