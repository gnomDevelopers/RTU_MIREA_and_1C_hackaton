package entities

type Class struct {
	Id                   int      `json:"id"`
	Name                 string   `json:"name"`
	AcademicDisciplineId int      `json:"academic_discipline_id"`
	GroupNames           []string `json:"group_names"`
	TeacherNames         []string `json:"teacher_names"`
	Type                 string   `json:"type"`
	AuditoryId           int      `json:"auditory_id"`
	Date                 string   `json:"date"`
	Weekday              int      `json:"weekday"`
	Week                 int      `json:"week"`
	TimeStart            string   `json:"time_start"`
	TimeEnd              string   `json:"time_end"`
}
