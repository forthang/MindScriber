package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Note    string
	OwnerId int
}
