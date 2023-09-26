package web

type UsersLogin struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"min=1"`
	Password string `json:"password"`
}

type UserChechk struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"min=1"`
	Password string `json:"password"`
}
