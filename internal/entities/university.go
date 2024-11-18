package entities

type University struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUniversityResponse struct {
	Id int `json:"id"`
}

type UpdateDeleteUniversityResponse struct {
	Message string `json:"message"`
}
