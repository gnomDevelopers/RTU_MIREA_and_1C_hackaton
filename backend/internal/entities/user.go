package entities

type User struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ID           int    `json:"id"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ID           int    `json:"id"`
}
