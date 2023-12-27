package dtos

type SignUpRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponseDto struct {
	Email string `json:"email"`
}