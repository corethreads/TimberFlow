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

func (s *UserRepository) CreateUser(user *entity.User) error {
	return s.DB.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
