package verification

type InputVerification struct {
	UserId int    `json:"user_id" validate:"required,numeric"`
	Types  string `json:"types" validate:"required,in=verification+forgot-password"`
	Code   string `json:"code" validate:"required,max=6"`
}

type updateStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

type InputVerifyForgoutPassword struct {
	Token           string `json:"token" validate:"required,size=20"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,password-custom"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
}