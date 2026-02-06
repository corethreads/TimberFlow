package service

import (
	"server/internal/auth/helpers"
	"server/internal/auth/models/dto"
	"server/internal/auth/models/entity"
	"server/internal/auth/repository"
	"server/internal/auth/types"
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
	//Todo generate password hash
	//Todo add user to Database
}
