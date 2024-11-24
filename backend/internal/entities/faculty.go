package entities

type Faculty struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	University string `json:"university"`
}

type CreateFacultyRequest struct {
	Name       string `json:"name"`
	University string `json:"university"`
}

type CreateFacultyResponse struct {
	ID int `json:"id"`
}

type GetFacultyRequest struct {
	UniversityName string `json:"university_name"`
}

type UpdateFacultyRequest struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	University string `json:"university"`
}
