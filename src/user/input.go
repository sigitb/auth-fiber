package user

type InputRegister struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email,min=6"`
	Password string `json:"password" validate:"required,min=8,max=14,password-custom"`
	RoleId   uint   `json:"role_id" validate:"required,numeric"`
}

type InputLogin struct {
	Email    string `json:"email" validate:"required,email,min=6"`
	Password string `json:"password" validate:"required,min=8,max=14"`
}

type InputVerification struct {
	Email string `json:"email" validate:"required,email,min=6"`
	Code  string `json:"code" validate:"required,max=6"`
}

type InputSendOtp struct {
	Email string `json:"email" validate:"required,email,min=6"`
	Types string `json:"types" validate:"required,in=verification+forgot-password"`
}