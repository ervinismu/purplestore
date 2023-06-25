package schema

type RegisterRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required,alphanum" json:"username"`
	Password string `validate:"required,min=8,alphanum" json:"password"`
}
