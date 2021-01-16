package project

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	OwnerID     int    `json:"creator"               db:"creator"`
	IsDefault   bool   `json:"isDefault"             gorm:"default:false"`
	Title       string `json:"title"                 gorm:"not null"`
	Description string `json:"description,omitempty" db:"description"`
	Logo        string `json:"logo,omitempty"        db:"logo"`
}
