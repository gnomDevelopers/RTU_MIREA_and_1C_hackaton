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

type CreateGradeResponse struct {
	Id int `json:"id"`
}

type UpdateDeleteGradeResponse struct {
	Message string `json:"message"`
}
