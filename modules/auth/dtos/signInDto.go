package dtos

type SignInRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignInResponseDto struct {
	Email string `json:"email"`
	Token string `json:"token"`
}