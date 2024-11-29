package repository

import (
	"context"
	"server/internal/entities"
)

type UserRepository interface {
	GetById(context.Context, int) (*entities.User, error)
	GetByEmail(context.Context, string) (*entities.User, error)
	GetByUniversity(context.Context, string) (*[]entities.User, error)
	Exists(context.Context, string) (bool, error)
	CreateUser(context.Context, *entities.User) (*entities.User, error)
	GetInfoById(context.Context, int) (*entities.UserInfo, error)
	UpdateRole(context.Context, int, string) error
	GetAllTeachers(context.Context, string) (*[]entities.User, error)
}

//type UserDataRepository interface {
//	Add(context.Context, *entities.UserData) error
//}

type UniversityRepository interface {
	Exists(context.Context, string) (bool, error)
	Create(context.Context, *entities.CreateUniversityRequest) (int, error)
	GetById(context.Context, int) (*entities.University, error)
	GetByName(context.Context, string) (*entities.University, error)
	GetByUserID(context.Context, int) (*entities.University, error)
	GetAll(context.Context) (*[]entities.University, error)
	Update(context.Context, int, string) error
	Delete(context.Context, int) error
}

type CampusRepository interface {
	Create(context.Context, *[]entities.Campus) ([]int, error)
	GetById(context.Context, int) (*entities.Campus, error)
	GetByAddress(context.Context, string) (*entities.Campus, error)
	GetByName(context.Context, string) (*entities.Campus, error)
	GetByUniversity(context.Context, string) (*[]entities.Campus, error)
	GetAll(context.Context) (*[]entities.Campus, error)
	Update(context.Context, *entities.Campus) error
	Delete(context.Context, int) error
}

type AudienceRepository interface {
	Create(context.Context, *[]entities.Audience) ([]int, error)
	GetById(context.Context, int) (*entities.Audience, error)
	GetByType(context.Context, string) (*[]entities.Audience, error)
	GetByCampus(context.Context, string) (*[]entities.Audience, error)
	GetByProfile(context.Context, string) (*[]entities.Audience, error)
	GetByCapacity(context.Context, int) (*[]entities.Audience, error)
	GetByUniversity(context.Context, string) (*[]entities.Audience, error)
	Update(context.Context, *entities.Audience) error
	Delete(context.Context, int) error
}

type ClassRepository interface {
	Create(context.Context, *[]entities.Class) ([]int, error)
	GetById(context.Context, int) (*entities.Class, error)
	GetByGroupName(context.Context, string) (*[]entities.Class, error)
	GetOptionals(context.Context, string, string) (*[]entities.Class, error)
	GetByNameAndGroup(context.Context, string, string) (*[]entities.Class, error)
	GetByNameAndGroupWithoutLk(context.Context, string, string) (*[]entities.Class, error)
	GetByTeacherName(context.Context, string, string) (*[]entities.Class, error)
	GetByAuditory(context.Context, string) (*[]entities.Class, error)
	SearchGroups(context.Context, string) ([]string, error)
	SearchTeachers(context.Context, string) ([]string, error)
	SearchNames(context.Context, string) ([]string, error)
	SearchNamesWithGroup(context.Context, string) ([]string, error)
	Update(context.Context, *entities.Class) error
	Delete(context.Context, int) error
	GetAllParticipants(context.Context, int) (*[]entities.ClassParticipant, error)
}

type GroupRepository interface {
	Exists(context.Context, string) (bool, error)
	Create(context.Context, *entities.CreateGroupRequest) (int, error)
	GetById(context.Context, int) (*entities.Group, error)
	GetByUserID(context.Context, int) (*entities.Group, error)
	GetByName(context.Context, string) (*entities.Group, error)
	GetAll(context.Context, string) (*[]entities.Group, error)
	GetGroupMembers(context.Context, string) (*[]entities.GroupMember, error)
}

type UserScheduleRepository interface {
	Create(context.Context, *entities.UserSchedule) (int, error)
	GetByUserId(context.Context, int) (*[]entities.UserSchedule, error)
	Update(context.Context, *entities.UserSchedule) error
	Delete(context.Context, int) error
}

type UserDataRepository interface {
	AddUser(context.Context, *entities.UserData) (int, error)
	AddAdmin(context.Context, *entities.UserData) (int, error)
	GetEducationalDirection(context.Context, int) (string, error)
	GetAll(context.Context, int) (*[]entities.GetUserDataResponse, error)
	GetAllByRole(context.Context, int, string) (*[]entities.GetUserDataResponse, error)
	GetById(context.Context, int) (*entities.UserData, error)
}

type FacultyRepository interface {
	Exists(context.Context, string) (bool, error)
	Create(context.Context, *entities.CreateFacultyRequest) (int, error)
	GetById(context.Context, int) (*entities.Faculty, error)
	GetByName(context.Context, string) (*entities.Faculty, error)
	GetAllByUniName(context.Context, *entities.GetFacultyRequest) (*[]entities.Faculty, error)
	GetAll(context.Context) (*[]entities.Faculty, error)
	Update(context.Context, *entities.UpdateFacultyRequest) error
	Delete(context.Context, int) error
}

type DepartmentRepository interface {
	Exists(context.Context, string) (bool, error)
	Create(context.Context, *entities.CreateDepartmentRequest) (int, error)
	GetById(context.Context, int) (*entities.Department, error)
	GetByName(context.Context, string) (*entities.Department, error)
	GetAll(context.Context) (*[]entities.Department, error)
	GetByUniversity(context.Context, string) (*[]entities.Department, error)
	Update(context.Context, *entities.UpdateDepartmentRequest) error
	Delete(context.Context, int) error
}

type GradeRepository interface {
	Create(context.Context, *entities.Grade) (int, error)
	GetByUserIdAndClassId(context.Context, int, int) (*entities.Grade, error)
}

type ScoreRepository interface {
	Update(context.Context, *entities.Score) error
	Get(context.Context, *entities.Score) (*entities.Score, error)
}

type AcademicDisciplineRepository interface {
	Create(context.Context, *entities.AcademicDiscipline) (int, error)
	GetByEducationalDirectionAndSemester(context.Context, string, int) (*[]entities.AcademicDiscipline, error)
	GetByEducationalDirectionAndName(context.Context, string, string) (*entities.AcademicDiscipline, error)
	Update(context.Context, *entities.AcademicDiscipline) error
	Delete(context.Context, int) error
	GetDisciplinesByGroupName(context.Context, string) (*[]entities.AcademicDiscipline, error)
}

type GpaRepository interface {
	Update(context.Context, int, float64) error
	Get(context.Context, int) (*entities.Gpa, error)
	GetByHighestGpa(context.Context, float64) (*[]entities.Gpa, error)
}

type WorkRepository interface {
	Exists(context.Context, int) (bool, error)
	ExistsHR(context.Context, string) (bool, error)
	Create(context.Context, *entities.WorkUser) error
	CreateHR(context.Context, *entities.HR) error
	GetByIdWorkUser(context.Context, int) (*entities.FullWorkUser, error)
	GetByIdWorkUserCVPath(context.Context, int) (string, error)
	GetByEmailHR(context.Context, string) (*entities.HR, error)
	UpdateWorkUser(context.Context, *entities.WorkUserUpdateRequest) error
	UpdateCVPath(context.Context, *entities.UpdateCV) error
	CreateResponse(context.Context, *entities.Response) error
	GetWorkUserResponses(context.Context, int) (*[]entities.Response, error)
	GetHRResponses(context.Context, int) (*[]entities.CandidateResponse, error)
	GetAllWorkUser(context.Context) (*[]entities.FullWorkUser, error)
}

type AchievementRepository interface {
	Create(context.Context, *entities.CreateAchievementRequest) (*entities.CreateAchievementResponse, error)
	GetById(context.Context, int) (*[]entities.Achievement, error)
	Delete(context.Context, int) error
}

type VisitingRepository interface {
	Exist(ctx context.Context, visiting *entities.Visiting) (bool, error)
	Create(ctx context.Context, visiting *entities.Visiting) (int, error)
	GetByUserIdAndClassId(ctx context.Context, userID, classID int) (*entities.Visiting, error)
	Update(ctx context.Context, visiting *entities.Visiting) error
	GetGroupVisiting(context.Context, int) (*[]entities.VisitingInfo, error)
}

type Repository struct {
	User                  UserRepository
	UserData              UserDataRepository
	University            UniversityRepository
	Campus                CampusRepository
	Group                 GroupRepository
	Audience              AudienceRepository
	AchievementRepository AchievementRepository
	Class                 ClassRepository
	Faculty               FacultyRepository
	Department            DepartmentRepository
	UserSchedule          UserScheduleRepository
	Grade                 GradeRepository
	Score                 ScoreRepository
	AcademicDiscipline    AcademicDisciplineRepository
	Gpa                   GpaRepository
	Work                  WorkRepository
	Visiting              VisitingRepository
}
