package models

type Session struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Token     string `gorm:"uniqueIndex;not null"`
	ExpiresAt int64  `gorm:"not null"`
}
