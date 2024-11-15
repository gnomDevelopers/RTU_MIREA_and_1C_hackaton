package entities

type Campus struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	UniversityId int    `json:"university_id"`
	Address      string `json:"address"`
}

type CreateCampusResponse struct {
	Id int `json:"id"`
}

type UpdateDeleteCampusResponse struct {
	Message string `json:"message"`
}
