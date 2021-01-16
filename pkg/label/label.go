package label

import "gorm.io/gorm"

type Label struct {
	Title       string `json:"title"                  gorm:"not null"`
	Description string `json:"description,omitempty"  db:"description"`
	Colour      string `json:"colour"                 db:"colour"`
	OwnerID     int    `json:"creator"                db:"creator"`
	gorm.Model
}

type Service interface {
}

type Repository interface {
}
