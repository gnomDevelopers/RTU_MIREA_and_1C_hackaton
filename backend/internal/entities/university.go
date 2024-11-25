package entities

type University struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Postfix string `json:"postfix"`
}

type CreateUniversityRequest struct {
	Name    string `json:"name"`
	Postfix string `json:"postfix"`
}

type CreateUniversityResponse struct {
	Id int `json:"id"`
}

type UpdateDeleteUniversityResponse struct {
	Message string `json:"message"`
}
