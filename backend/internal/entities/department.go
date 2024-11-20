package entities

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateDepartmentRequest struct {
	Name string `json:"name"`
}

type CreateDepartmentResponse struct {
	ID int `json:"id"`
}

type UpdateDepartmentRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
