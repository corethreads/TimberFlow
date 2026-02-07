package controller

import (
	"net/http"
	"server/internal/auth/models/dto"
	"server/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type authController struct {
	userService *service.AuthService
}

func NewAuthController(authservice *service.AuthService) *authController {
	return &authController{userService: authservice}
}

//http communication Handler

// TODO handler for Register
func (ac *authController) Register(c *gin.Context) {
	//TODO Request Going
	var request dto.RequestDTO
	//TODO Bind Json Request to understable format
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	//TODO Create user using Service
	user, err := ac.userService.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	//TODO return response from handler
	//
	c.JSON(http.StatusOK, user)
}
