package entities

type Class struct {
	Id            int      `json:"id"`
	Name          string   `json:"name"`
	GroupNames    []string `json:"group_names"`
	TeacherNames  []string `json:"teacher_names"`
	Type          string   `json:"type"`
	Auditory      string   `json:"auditory"`
	Date          string   `json:"date"`
	Weekday       int      `json:"weekday"`
	Week          int      `json:"week"`
	TimeStart     string   `json:"time_start"`
	TimeEnd       string   `json:"time_end"`
	UniversityStr string   `json:"university"`
}

type CreateClassesRequest struct {
	Name         string   `json:"name"`
	GroupNames   []string `json:"group_names"`
	TeacherNames []string `json:"teacher_names"`
	Type         string   `json:"type"`
	Auditory     string   `json:"auditory"`
	Date         string   `json:"date"`
	Weekday      int      `json:"weekday"`
	Week         int      `json:"week"`
	TimeStart    string   `json:"time_start"`
	TimeEnd      string   `json:"time_end"`
}

type CreateClassesResponse struct {
	Id int `json:"id"`
}

type UpdateDeleteClassResponse struct {
	Message string `json:"message"`
}

type ParseScheduleResponse struct {
	Message string `json:"message"`
}

type ScheduleGroups struct {
	Groups []string `json:"groups"`
}

type ScheduleTeachers struct {
	Teachers []string `json:"teachers"`
}

type ScheduleNames struct {
	Names []string `json:"names"`
}

type GradeClass struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Date   string  `json:"date"`
	Grades []Grade `json:"grades"`
}
