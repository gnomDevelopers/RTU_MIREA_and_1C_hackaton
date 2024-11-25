package entities

type User struct {
	ID                   int    `json:"id"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	LastName             string `json:"last_name"`
	FirstName            string `json:"first_name"`
	FatherName           string `json:"father_name"`
	UniversityID         int    `json:"university_id"`
	Role                 string `json:"role"`
	FacultyID            int    `json:"faculty_id"`
	Group                string `json:"group"`
	DepartmentID         int    `json:"department_id"`
	EducationalDirection string `json:"educational_direction"`
	IsDeleted            bool   `json:"is_deleted"`
	IsPasswordChanged    bool   `json:"is_password_changed"`
	CreatedAt            string `json:"created_at"`
}

type CreateUserRequest struct {
	Password             string `json:"password"`
	Email                string `json:"email"`
	LastName             string `json:"last_name"`
	FirstName            string `json:"first_name"`
	FatherName           string `json:"father_name"`
	UniversityID         int    `json:"university_id"`
	Role                 string `json:"role"`
	FacultyID            int    `json:"faculty_id"`
	Group                string `json:"group"`
	DepartmentID         int    `json:"department_id"`
	EducationalDirection string `json:"educational_direction"`
}

type CreateUserResponse struct {
	ID         int    `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	FatherName string `json:"father_name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ID           int    `json:"id"`
}

type GroupMember struct {
	ID         int    `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	FatherName string `json:"father_name"`
}
