package models

type Note struct {
	ID      uint   `gorm:"primaryKey"`
	Note    string `gorm:"uniqueIndex;not null"`
	OwnerID uint   `gorm:"not null"`
}
