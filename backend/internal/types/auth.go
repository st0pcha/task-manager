package types

type AuthDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterDTO struct {
	AuthDTO

	PasswordRepeat string `json:"repeated_password"`
}
