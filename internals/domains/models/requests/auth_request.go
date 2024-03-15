package requests

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email" example:"example@gmail.com"`
	UserName string `json:"username" validate:"required" example:"user"`
	Password string `json:"password" validate:"required,min=8,max=20" example:"12345678"`
	Age      int    `json:"age" validate:"gte=0,lte=130" example:"25"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
