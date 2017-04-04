package models

import "time"

// Model of go-blog
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
