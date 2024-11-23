package entities

type UserSchedule struct {
	Id         int    `json:"id"`
	UserDataId int    `json:"user_data_id"`
	Name       string `json:"name"`
	Date       string `json:"date"`
	TimeStart  string `json:"time_start"`
	TimeEnd    string `json:"time_end"`
}

type CreateUserScheduleRequest struct {
	Name      string `json:"name"`
	Date      string `json:"date"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
}

type GetUserScheduleResponse struct {
	UserSchedule *[]UserSchedule `json:"user_schedule"`
	Classes      *[]Class        `json:"classes"`
}

type CreateUserScheduleResponse struct {
	Id int `json:"id"`
}

type UpdateUserScheduleRequest struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
}

type UpdateDeleteUserScheduleResponse struct {
	Message string `json:"message"`
}
