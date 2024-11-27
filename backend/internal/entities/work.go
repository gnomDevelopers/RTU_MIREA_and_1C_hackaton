package entities

type WorkUser struct {
	Id          int      `json:"id"`
	PhoneNumber string   `json:"phone_number"`
	Telegram    string   `json:"telegram"`
	Skills      []string `json:"skills"`
	CVPath      string   `json:"cv_path"`
	HideProfile bool     `json:"hide_profile"`
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
