package dtos

type SignUpRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignUpResponseDto struct {
	Email string `json:"email"`
}
