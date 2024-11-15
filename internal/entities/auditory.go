package entities

type Auditory struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	CampusId int    `json:"campus_id"`
	Type     string `json:"type"`
	Profile  string `json:"profile"`
	Capacity int    `json:"capacity"`
}
