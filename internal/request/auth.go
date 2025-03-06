package request

type RegisterUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,max=200,min=1" json:"password"`
}
