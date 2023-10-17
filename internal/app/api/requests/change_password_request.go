package requests

type ChangePasswordRequest struct {
	Code     string `json:"code" validate:"required,lte=5,gte=5"`
	Password string `json:"password" validate:"required,gte=6"`
	Email    string `json:"email" validate:"required,email,exists=users"`
}
