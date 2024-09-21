package storage

import (
	"github.com/forthang/esketit/internal/config"
	"github.com/forthang/esketit/models"
	"gorm.io/gorm"
)

type Storage struct {
	Postgres *gorm.DB
}

func NewStorage(cfg config.Config) *Storage {
	db := InitDB(cfg)
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Note{})
	return &Storage{
		Postgres: db,
	}
}

func (s *Storage) CreateUser(user models.User) error {
	return s.Postgres.Create(&user).Error
}

func (s *Storage) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := s.Postgres.First(&user, id).Error
	return user, err
}

func (s *Storage) UpdateUser(user models.User) error {
	return s.Postgres.Save(&user).Error
}

func (s *Storage) DeleteUser(id uint) error {
	return s.Postgres.Delete(&models.User{}, id).Error
}
