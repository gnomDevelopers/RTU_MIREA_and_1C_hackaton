package entities

type Faculty struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateFacultyRequest struct {
	Name string `json:"name"`
}

type CreateFacultyResponse struct {
	ID int `json:"id"`
}

type UpdateFacultyRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
