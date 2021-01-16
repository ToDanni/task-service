package project

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Creator     int    `json:"creator"               db:"creator"`
	IsDefault   bool   `json:"isDefault"             db:"is_default"`
	Title       string `json:"title"                 db:"title"`
	Description string `json:"description,omitempty" db:"description"`
	Logo        string `json:"logo,omitempty"        db:"logo"`
}
