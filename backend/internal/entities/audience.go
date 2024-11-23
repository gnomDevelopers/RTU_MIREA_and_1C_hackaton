package entities

type Audience struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Campus   string `json:"campus"`
	Type     string `json:"type"`
	Profile  string `json:"profile"`
	Capacity int    `json:"capacity"`
}

type CreateAudiencesRequest struct {
	Name     string `json:"name"`
	Campus   string `json:"campus"`
	Type     string `json:"type"`
	Profile  string `json:"profile"`
	Capacity int    `json:"capacity"`
}

type CreateAudiencesResponse struct {
	Id int `json:"id"`
}

type UpdateDeleteAudienceResponse struct {
	Message string `json:"message"`
}
