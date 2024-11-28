package entities

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateGroupRequest struct {
	Name string `json:"name"`
}

type CreateGroupResponse struct {
	ID int `json:"id"`
}

type GetGroupRequest struct {
	Name string `json:"name"`
}
