package helpers

import (
	"server/internal/auth/models/dto"
)

//check if user wrote credentials yes/no

func HasRequiredFields(user *dto.RequestDTO) bool {
	return user.BusinessName != "" && user.Username != "" && user.Email != "" && user.Password != ""
}
