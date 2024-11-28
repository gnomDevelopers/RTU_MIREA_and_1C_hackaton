package entities

type WorkUser struct {
	Id          int      `json:"id"`
	PhoneNumber string   `json:"phone_number"`
	Telegram    string   `json:"telegram"`
	Skills      []string `json:"skills"`
	CVPath      string   `json:"cv_path"`
	HideProfile bool     `json:"hide_profile"`
}

type WorkUserUpdateRequest struct {
	Id          int      `json:"id"`
	PhoneNumber string   `json:"phone_number"`
	Telegram    string   `json:"telegram"`
	Skills      []string `json:"skills"`
	CVPath      string   `json:"cv_path"`
}

type WorkUserUpdateResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Id         int `json:"id"`
	HRId       int `json:"hr_id"`
	WorkUserId int `json:"work_user_id"`
}

type CreateResponse struct {
	HRId       int `json:"hr_id"`
	WorkUserId int `json:"work_user_id"`
}

type HR struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	FatherName string `json:"father_name"`
	Company    string `json:"company"`
	ImageLink  string `json:"image_link"`
}

type LoginWorkUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ID           int    `json:"id"`
}

type ExistsResponse struct {
	Exists bool `json:"exists"`
}

type FullWorkUser struct {
	Id          int      `json:"id"`
	LastName    string   `json:"last_name"`
	FirstName   string   `json:"first_name"`
	FatherName  string   `json:"father_name"`
	University  string   `json:"university"`
	Gpa         float64  `json:"gpa"`
	PhoneNumber string   `json:"phone_number"`
	Telegram    string   `json:"telegram"`
	Skills      []string `json:"skills"`
	CVPath      string   `json:"cv_path"`
}
