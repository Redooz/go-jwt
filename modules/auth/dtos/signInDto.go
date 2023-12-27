package dtos

type SignInRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponseDto struct {
	Email string `json:"email"`
}