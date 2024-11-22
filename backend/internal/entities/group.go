package entities

type Group struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

type CreateGroupRequest struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}
