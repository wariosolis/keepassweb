package models

import "time"

// Passwd model
type Passwd struct {
	ID             string `sql:"type:varchar(36)"`
	User           User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID         int
	Category       Category `gorm:"ForeignKey:CategoryID;AssociationForeignKey:ID"`
	CategoryID     int
	Title          string `gorm:"size:255"`
	Username       string `gorm:"size:255"`
	URL            string `gorm:"size:255"`
	Notes          string `gorm:"size:255"`
	ExpirationDate *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	HashedPassword []byte
}
