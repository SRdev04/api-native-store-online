package web

type UsersToken struct {
	Id            int    `json:"id"`
	Access_Token  string `json:"access_token" validate:"min=1"`
	Refresh_Token string `json:"refresh_token"`
	Token_Type    string `json:"token_type"`
}

type UsersRespon struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `validate:"min=1"`
	Password string `json:"password"`
}
