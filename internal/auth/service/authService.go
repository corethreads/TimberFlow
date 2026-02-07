package service

import (
	"errors"
	"os"
	"server/internal/auth/helpers"
	"server/internal/auth/models/dto"
	"server/internal/auth/models/entity"
	"server/internal/auth/repository"
	"server/internal/auth/types"
	"server/internal/auth/utils"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	useRepo *repository.UserRepository
}

func NewauthService(userepo *repository.UserRepository) *AuthService {
	return &AuthService{useRepo: userepo}
}

func (s *AuthService) CreateUser(request dto.RequestDTO) (*dto.ResponseDTO, error) {
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
	//Todo add user to Storage Database
	user := &entity.User{
		ID:            uuid.New().String(),
		Business_ID:   uuid.New().String(),
		Business_Name: request.BusinessName,
		Username:      request.Username,
		Email:         request.Email,
		Password:      hashedPassword,
	}

	if err := s.useRepo.CreateUser(user); err != nil {
		return nil, errors.New("User Not created")
	}

	//Generate JWT for user for further usage
	secretkey := os.Getenv("JWT_SECRET")
	expiry := 3 * time.Hour

	jwtToken, err := utils.Generatejwt(secretkey, user.ID, user.Email, expiry)
	if err != nil {
		return nil, err
	}

	//Generate Response for User
	response := &dto.ResponseDTO{
		Token: jwtToken,
		User: dto.Response{
			ID:            user.ID,
			BusinessID:    user.Business_ID,
			Business_name: user.Business_Name,
			Username:      user.Username,
			Email:         user.Email,
		},
	}

	return response, nil

}
