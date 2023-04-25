package requests

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validte:"required,gte-6,lte=255"`
}
