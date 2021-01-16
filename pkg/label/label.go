package label

import (
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	Creator     int    `json:"creator"               db:"creator"`
	Title       string `json:"title"                 db:"title"`
	Description string `json:"description,omitempty" db:"description"`
	Colour      string `json:"colour"                 db:"colour"`
}
