package entities

type AcademicDiscipline struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	EducationalDirection string `json:"educational_direction"`
	Semester             int    `json:"semester"`
	LectureHours         int    `json:"lecture_hours"`
	PracticeHours        int    `json:"practice_hours"`
	LabHours             int    `json:"lab_hours"`
	IndividualHours      int    `json:"individual_hours"`
	TypeOfAssessment     string `json:"type_of_assessment"`
}

type CreateAcademicDisciplineRequest struct {
	Name                 string `json:"name"`
	EducationalDirection string `json:"educational_direction"`
	Semester             int    `json:"semester"`
	LectureHours         int    `json:"lecture_hours"`
	PracticeHours        int    `json:"practice_hours"`
	LabHours             int    `json:"lab_hours"`
	IndividualHours      int    `json:"individual_hours"`
	TypeOfAssessment     string `json:"type_of_assessment"`
}

type CreateAcademicDisciplineResponse struct {
	Id int `json:"id"`
}
