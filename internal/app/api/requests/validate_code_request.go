package requests

type ValidateCodeRequest struct {
	Code  string `json:"code" validate:"required,lte=5,gte=5"`
	Email string `json:"email" validate:"required,email,exists=users"`
}
