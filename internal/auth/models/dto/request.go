package dto

type RequestDTO struct {
	BusinessName string `json:"businessname" binding:"required"`
	Username     string `json:"Username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
}
