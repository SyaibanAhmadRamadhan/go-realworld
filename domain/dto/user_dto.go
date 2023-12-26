package dto

type RequestLoginUser struct {
	EmailOrUsername string `json:"email_or_username"`
	Password        string `json:"password"`
}

type RequestRegisterUser struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

type RequestUpdateUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Bio      string `json:"bio"`
}

type ResponseUser struct {
	Id       string `json:"id"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Bio      string `json:"bio,omitempty"`
}

type ResponseToken struct {
	AccessToken string `json:"access_token"`
}
