package entities

type Score struct {
	UserId      int    `json:"user_id"`
	Sum         int    `json:"sum"`
	Count       int    `json:"count"`
	SubjectName string `json:"subject_name"`
}
