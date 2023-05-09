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
