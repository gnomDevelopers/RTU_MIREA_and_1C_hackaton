package entities

type Campus struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	University string `json:"university"`
	Address    string `json:"address"`
}

type CreateCampusesRequest struct {
	Name       string `json:"name"`
	University string `json:"university"`
	Address    string `json:"address"`
}

type CreateCampusesResponse struct {
	Id int `json:"id"`
}

type UpdateDeleteCampusResponse struct {
	Message string `json:"message"`
}
