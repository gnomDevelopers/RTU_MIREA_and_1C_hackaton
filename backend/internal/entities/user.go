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
	GroupID              int    `json:"group_id"`
	DepartmentID         int    `json:"department_id"`
	EducationalDirection string `json:"educational_direction"`
	IsDeleted            bool   `json:"is_deleted"`
	IsPasswordChanged    bool   `json:"is_password_changed"`
	CreatedAt            string `json:"created_at"`
}

type UserInfo struct {
	ID                   int    `json:"id"`
	Email                string `json:"email"`
	LastName             string `json:"last_name"`
	FirstName            string `json:"first_name"`
	FatherName           string `json:"father_name"`
	UniversityName       string `json:"university_name"`
	Role                 string `json:"role"`
	FacultyID            int    `json:"faculty_id"`
	GroupID              int    `json:"group_id"`
	DepartmentID         int    `json:"department_id"`
	EducationalDirection string `json:"educational_direction"`
}

type CreateUserRequest struct {
	Password             string `json:"password"`
	LastName             string `json:"last_name"`
	FirstName            string `json:"first_name"`
	FatherName           string `json:"father_name"`
	UniversityID         int    `json:"university_id"`
	Role                 string `json:"role"`
	FacultyID            int    `json:"faculty_id,omitempty"`
	GroupID              int    `json:"group_id,omitempty"`
	DepartmentID         int    `json:"department_id,omitempty"`
	EducationalDirection string `json:"educational_direction,omitempty"`
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
	Group      string `json:"group"`
}

type UpdateRoleRequest struct {
	NewRole string `json:"new_role"`
}

type ClassParticipant struct {
	ID         int    `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	FatherName string `json:"father_name"`
}
