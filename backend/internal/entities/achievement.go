package entities

type Achievement struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Result string `json:"result"`
	UserID int    `json:"user_id"`
}

type CreateAchievementRequest struct {
	Name   string `json:"name"`
	Result string `json:"result"`
	UserID int    `json:"user_id"`
}

type CreateAchievementResponse struct {
	ID int `json:"id"`
}
