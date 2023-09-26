package web

type UsersRegis struct {
	Username string `validate:"min=1"`
	Email    string `validate:"min=1"`
	Password string `json:"password"`
}
