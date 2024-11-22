package entities

//CREATE TABLE IF NOT EXISTS user_data (
//    id INTEGER PRIMARY KEY REFERENCES "user"(id),
//    last_name VARCHAR,
//    first_name VARCHAR,
//    father_name VARCHAR,
//    university_id INTEGER REFERENCES university(id),
//    permission_id INTEGER,
//    faculty_id INTEGER REFERENCES faculty(id),
//    department_id INTEGER REFERENCES department(id),
//    is_deleted BOOLEAN,
//    is_password_changed BOOLEAN,
//    created_at TIMESTAMP
//);

type UserData struct {
	ID                   int    `json:"id"`
	LastName             string `json:"last_name"`
	FirstName            string `json:"first_name"`
	FatherName           string `json:"father_name"`
	UniversityID         int    `json:"university_id"`
	PermissionID         int    `json:"permission_id"`
	FacultyID            int    `json:"faculty_id"`
	Group                string `json:"group"`
	DepartmentID         int    `json:"department_id"`
	EducationalDirection string `json:"educational_direction"`
	IsDeleted            bool   `json:"is_deleted"`
	IsPasswordChanged    bool   `json:"is_password_changed"`
	CreatedAt            string `json:"created_at"`
}

type AddUserDataRequest struct {
	ID                   int    `json:"id"`
	LastName             string `json:"last_name"`
	FirstName            string `json:"first_name"`
	FatherName           string `json:"father_name"`
	University           string `json:"university"`
	PermissionID         int    `json:"permission_id"`
	Faculty              string `json:"faculty"`
	Group                string `json:"group"`
	EducationalDirection string `json:"educational_direction"`
	Department           string `json:"department"`
}
