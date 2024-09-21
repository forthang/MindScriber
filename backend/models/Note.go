package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Data    string
	OwnerID int
}
