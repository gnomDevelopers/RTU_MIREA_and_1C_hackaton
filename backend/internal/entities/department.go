package entities

type Department struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	University string `json:"university"`
}

type CreateDepartmentRequest struct {
	Name       string `json:"name"`
	University string `json:"university"`
}

type CreateDepartmentResponse struct {
	ID int `json:"id"`
}

type UpdateDepartmentRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
