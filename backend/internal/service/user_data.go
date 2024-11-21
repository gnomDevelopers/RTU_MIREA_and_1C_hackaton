package service

import "server/internal/repository"

type UserData struct {
	UserRepository *repository.UserRepository
	//UserDataRepository *repository.UserDataRepository
	FacultyRepository    *repository.FacultyRepository
	DepartmentRepository *repository.DepartmentRepository
	UniversityRepository *repository.UniversityRepository
	UniversityCache      map[string]int
}

func NewUserData() *UserData {
	return &UserData{}
}
