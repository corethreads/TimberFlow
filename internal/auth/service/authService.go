package service

import (
	"errors"
	"server/internal/auth/helpers"
	"server/internal/auth/models/dto"
	"server/internal/auth/models/entity"
	"server/internal/auth/repository"
	"server/internal/auth/types"
	"server/internal/auth/utils"

	"github.com/google/uuid"
)

type authService struct {
	useRepo *repository.UserRepository
}

func NewauthService(userepo *repository.UserRepository) *authService {
	return &authService{useRepo: userepo}
}

func (s *authService) createUser(request dto.RequestDTO) (*entity.User, error) {
	//TODO check if user typed requests
	if !helpers.HasRequiredFields(&request) {
		return nil, types.NothinginFields
	}
	//TODO check if user exists
	if _, err := s.useRepo.GetUserByEmail(request.Email); err == nil {
		return nil, types.AlreadyAdded

	}
	//Todo generate password hash
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, errors.New("Hash Not created")
	}
	//Todo add user to Database
	user := &entity.User{
		Business_ID:   uuid.New().String(),
		Business_Name: request.BusinessName,
		Username:      request.Username,
		Email:         request.Email,
		Password:      hashedPassword,
	}

	if err := s.useRepo.CreateUser(user); err == nil {
		return nil, errors.New("User Not created")
	}

}
