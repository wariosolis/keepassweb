package models

// Category model
type Category struct {
	Model
	Slug        string `gorm:"type:varchar(75);unique_index"`
	Name        string
	Description string
}
