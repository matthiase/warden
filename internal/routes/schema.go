package routes

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateUserResponse struct {
	*User `json:"user"`
	Otp   string `json:"otp"`
}

type ConfirmUserRequest struct {
	Id  int    `json:"id"`
	Otp string `json:"otp"`
}

type ConfirmUserResponse struct {
	*User       `json:"user"`
	AccessToken string `json:"access_token"`
}
