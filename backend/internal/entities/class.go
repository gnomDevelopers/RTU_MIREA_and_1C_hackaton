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

type ScheduleGroup struct {
	Group string `json:"group"`
}

type ScheduleTeachers struct {
	Teacher string `json:"teacher"`
}

type ScheduleName struct {
	Name string `json:"name"`
}

type GetGradesBySubject struct {
	GroupMember []GroupMember `json:"group_member"`
	GradeClass  []GradeClass  `json:"grade_class"`
	UsersScore  []UsersScore  `json:"users_score"`
}

type GradeClass struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Date   string  `json:"date"`
	Grades []Grade `json:"grades"`
}
