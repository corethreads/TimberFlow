package repository

import (
	"server/internal/auth/models/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{DB: database}
}

func (s *UserRepository) createUser(user *entity.User) error {
	return s.DB.Create(user).Error
}
