package user

type RequestUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
