package models

// Category model
type Category struct {
	Model
	User        User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID      int
	Slug        string `gorm:"type:varchar(75);unique_index"`
	Name        string
	Description string
	Passwd      []Passwd
}
