package entities

//CREATE TABLE IF NOT EXISTS visiting (
//    id SERIAL PRIMARY KEY,
//    class_id INTEGER REFERENCES class(id),
//    user_id INTEGER REFERENCES users(id),
//    type VARCHAR NOT NULL
//);

type Visiting struct {
	ID      int    `json:"id"`
	ClassID int    `json:"class_id"`
	UserID  int    `json:"user_id"`
	Type    string `json:"type"`
}

type VisitingInfo struct {
	UserID int    `json:"user_id"`
	Type   string `json:"type"`
}

type CreateVisitingRequest struct {
	ClassID int    `json:"class_id"`
	UserID  int    `json:"user_id"`
	Type    string `json:"type"`
}

type CreateVisitingResponse struct {
	ID int `json:"id"`
}

type CheckInRequest struct {
	ID      int `json:"id"`
	ClassID int `json:"class_id"`
	UserID  int `json:"user_id"`
}
