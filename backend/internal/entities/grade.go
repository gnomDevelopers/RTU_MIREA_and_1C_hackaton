package entities

type Grade struct {
	Id      int `json:"id"`
	ClassId int `json:"class_id"`
	UserId  int `json:"user_id"`
	Value   int `json:"value"`
}

type CreateGradeRequest struct {
	ClassId int `json:"class_id"`
	UserId  int `json:"user_id"`
	Value   int `json:"value"`
}

type UsersScore struct {
	UserId       int     `json:"user_id"`
	AverageScore float64 `json:"average_score"`
	SumScore     int     `json:"sum_score"`
}

type CreateGradeResponse struct {
	Id int `json:"id"`
}

type UpdateDeleteGradeResponse struct {
	Message string `json:"message"`
}
