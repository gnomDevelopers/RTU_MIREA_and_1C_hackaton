package service

import "server/internal/repository"

type UserData struct {
	UserRepository *repository.UserRepository
	//UserDataRepository *repository.UserDataRepository
	FacultyRepository    *repository.FacultyRepository
	DepartmentRepository *repository.DepartmentRepository
	UniversityRepository *repository.UniversityRepository
	UniversityCache      map[string]int
	DepartmentCache      map[string]int
	FacultyCache         map[string]int
}

//func NewUserData(userRepo *repository.UserRepository, userDataRepo *repository.UserD) *UserData {
//	return &UserData{
//
//	}
//}
