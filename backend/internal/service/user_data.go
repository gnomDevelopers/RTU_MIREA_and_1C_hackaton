package service

import (
	"context"
	"fmt"
	"server/internal/entities"
	"server/internal/repository"
	"server/util"
	"time"
)

type UserDataService struct {
	UserRepository       repository.UserRepository
	UserDataRepository   repository.UserDataRepository
	FacultyRepository    repository.FacultyRepository
	DepartmentRepository repository.DepartmentRepository
	UniversityRepository repository.UniversityRepository
	UniversityCache      map[string]int
	DepartmentCache      map[string]int
	FacultyCache         map[string]int
	timeout              time.Duration
}

func NewUserDataService(userRepo repository.UserRepository, userDataRepo repository.UserDataRepository, facultyRepo repository.FacultyRepository,
	departmentRepo repository.DepartmentRepository, universityRepo repository.UniversityRepository) *UserDataService {
	return &UserDataService{
		UserRepository:       userRepo,
		UserDataRepository:   userDataRepo,
		FacultyRepository:    facultyRepo,
		DepartmentRepository: departmentRepo,
		UniversityRepository: universityRepo,
		UniversityCache:      make(map[string]int),
		DepartmentCache:      make(map[string]int),
		FacultyCache:         make(map[string]int),
		timeout:              time.Duration(10) * time.Second,
	}
}

func (s *UserDataService) Add(c context.Context, requests *[]entities.AddUserDataRequest) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	for _, request := range *requests {
		fullName := fmt.Sprintf("%s %s %s", request.LastName, request.FirstName, request.FatherName)

		// Создаём пользователя
		newUser := &entities.User{
			Login:    util.GenerateLogin(fullName),
			Password: util.GenerateTemporaryPassword(15),
		}
		user, err := s.UserRepository.CreateUser(ctx, newUser)
		if err != nil {
			return err
		}

		userData := &entities.UserData{
			ID:         user.ID,
			LastName:   request.LastName,
			FirstName:  request.FirstName,
			FatherName: request.FatherName,
		}

		// Университет
		universityID, ok := s.UniversityCache[request.University]
		if !ok {
			// Проверяем существование записи
			exists, err := s.UniversityRepository.Exists(ctx, request.University)
			if err != nil {
				return err
			}

			if exists {
				// Если существует, получаем данные
				university, err := s.UniversityRepository.GetByName(ctx, request.University)
				if err != nil {
					return err
				}
				universityID = university.Id
			} else {
				createUniversity := &entities.CreateUniversityRequest{
					Name: request.University,
				}
				universityID, err = s.UniversityRepository.Create(ctx, createUniversity)
				if err != nil {
					return err
				}
			}

			// Сохраняем ID в кэше
			s.UniversityCache[request.University] = universityID
		}
		userData.UniversityID = universityID

		// Факультет
		facultyID, ok := s.FacultyCache[request.Faculty]
		if !ok {
			exists, err := s.FacultyRepository.Exists(ctx, request.Faculty)
			if err != nil {
				return err
			}

			if exists {
				faculty, err := s.FacultyRepository.GetByName(ctx, request.Faculty)
				if err != nil {
					return err
				}
				facultyID = faculty.ID
			} else {
				createFaculty := &entities.CreateFacultyRequest{
					Name: request.Faculty,
				}
				facultyID, err = s.FacultyRepository.Create(ctx, createFaculty)
				if err != nil {
					return err
				}
			}

			s.FacultyCache[request.Faculty] = facultyID
		}
		userData.FacultyID = facultyID

		// Департамент
		departmentID, ok := s.DepartmentCache[request.Department]
		if !ok {
			exists, err := s.DepartmentRepository.Exists(ctx, request.Department)
			if err != nil {
				return err
			}

			if exists {
				department, err := s.DepartmentRepository.GetByName(ctx, request.Department)
				if err != nil {
					return err
				}
				departmentID = department.ID
			} else {
				newDepartment := &entities.CreateDepartmentRequest{
					Name: request.Department,
				}
				departmentID, err = s.DepartmentRepository.Create(ctx, newDepartment)
				if err != nil {
					return err
				}
			}

			s.DepartmentCache[request.Department] = departmentID
		}
		userData.DepartmentID = departmentID

	}

	return nil
}
