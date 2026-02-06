package dto

type Response struct {
	ID            string `json:"ID"`
	Business_name string `json:"business_name" binding:"required"`
	Username      string `json:"username"`
	Email         string `json:"email" binding:"required,email"`
}

type ResponseDTO struct {
	Token string   `json:"token"`
	User  Response `json:"user"`
}
