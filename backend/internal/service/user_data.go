package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	GroupRepository      repository.GroupRepository
	UniversityCache      map[string]int
	DepartmentCache      map[string]int
	FacultyCache         map[string]int
	GroupCache           map[string]int
	timeout              time.Duration
}

func NewUserDataService(userRepo repository.UserRepository, userDataRepo repository.UserDataRepository, facultyRepo repository.FacultyRepository,
	departmentRepo repository.DepartmentRepository, universityRepo repository.UniversityRepository, groupRepo repository.GroupRepository) *UserDataService {
	return &UserDataService{
		UserRepository:       userRepo,
		UserDataRepository:   userDataRepo,
		FacultyRepository:    facultyRepo,
		DepartmentRepository: departmentRepo,
		UniversityRepository: universityRepo,
		GroupRepository:      groupRepo,
		UniversityCache:      make(map[string]int),
		DepartmentCache:      make(map[string]int),
		FacultyCache:         make(map[string]int),
		GroupCache:           make(map[string]int),
		timeout:              time.Duration(30) * time.Second,
	}
}

func (s *UserDataService) Add(c context.Context, requests *[]entities.AddUserDataRequest) (*[]entities.AddUserDataResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	var responses []entities.AddUserDataResponse
	for _, request := range *requests {
		fullName := fmt.Sprintf("%s %s %s", request.LastName, request.FirstName, request.FatherName)
		generatedEmail := util.GenerateLogin(fullName)
		password := util.GenerateTemporaryPassword(15)
		hashedPassword, err := util.HashPassword(password)
		if err != nil {
			return nil, err
		}

		// Создаём пользователя
		newUser := &entities.User{
			Email:    generatedEmail,
			Password: hashedPassword,
		}

		ok, err := s.UserRepository.Exists(ctx, newUser.Email)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if ok {
			continue
		}

		user, err := s.UserRepository.CreateUser(ctx, newUser)
		if err != nil {
			return nil, err
		}

		userData := &entities.UserData{
			ID:                   user.ID,
			LastName:             request.LastName,
			FirstName:            request.FirstName,
			FatherName:           request.FatherName,
			EducationalDirection: request.EducationalDirection,
		}

		// Университет
		universityID, ok := s.UniversityCache[request.University]
		if !ok {
			universityID, err = s.getOrCreateUniversity(ctx, request.University)
			if err != nil {
				return nil, err
			}
			s.UniversityCache[request.University] = universityID
		}
		userData.UniversityID = universityID

		// Факультет
		facultyID, ok := s.FacultyCache[request.Faculty]
		if !ok {
			facultyID, err = s.getOrCreateFaculty(ctx, request.Faculty)
			if err != nil {
				return nil, err
			}
			s.FacultyCache[request.Faculty] = facultyID
		}
		userData.FacultyID = facultyID

		// Департамент
		departmentID, ok := s.DepartmentCache[request.Department]
		if !ok {
			departmentID, err = s.getOrCreateDepartment(ctx, request.Department)
			if err != nil {
				return nil, err
			}
			s.DepartmentCache[request.Department] = departmentID
		}
		userData.DepartmentID = departmentID

		// Группа
		addStudent := &entities.CreateGroupRequest{
			Name:   request.Group,
			UserID: user.ID,
		}
		userData.Group = request.Group
		_, err = s.GroupRepository.Create(ctx, addStudent)
		if err != nil {
			return nil, err
		}

		_, err = s.UserDataRepository.AddStudent(ctx, userData)
		if err != nil {
			return nil, err
		}

		responses = append(responses, entities.AddUserDataResponse{
			LastName:   request.LastName,
			FirstName:  request.FirstName,
			FatherName: request.FatherName,
			Email:      generatedEmail,
			Password:   password,
		})
	}

	return &responses, nil
}

// getOrCreateUniversity обрабатывает добавление или получение университета
func (s *UserDataService) getOrCreateUniversity(ctx context.Context, name string) (int, error) {
	exists, err := s.UniversityRepository.Exists(ctx, name)
	if err != nil && err != sql.ErrNoRows { // Пропускаем ошибку, если запись не найдена
		return 0, err
	}

	if exists {
		university, err := s.UniversityRepository.GetByName(ctx, name)
		if err != nil {
			return 0, err
		}
		return university.Id, nil
	}

	createUniversity := &entities.CreateUniversityRequest{
		Name: name,
	}
	return s.UniversityRepository.Create(ctx, createUniversity)
}

// getOrCreateFaculty обрабатывает добавление или получение факультета
func (s *UserDataService) getOrCreateFaculty(ctx context.Context, name string) (int, error) {
	log.Println("faculty exists")
	exists, err := s.FacultyRepository.Exists(ctx, name)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if exists {
		log.Println("faculty get by name")
		faculty, err := s.FacultyRepository.GetByName(ctx, name)
		if err != nil {
			return 0, err
		}
		return faculty.ID, nil
	}

	createFaculty := &entities.CreateFacultyRequest{
		Name: name,
	}
	log.Println("faculty create")
	return s.FacultyRepository.Create(ctx, createFaculty)
}

// getOrCreateDepartment обрабатывает добавление или получение департамента
func (s *UserDataService) getOrCreateDepartment(ctx context.Context, name string) (int, error) {
	exists, err := s.DepartmentRepository.Exists(ctx, name)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if exists {
		department, err := s.DepartmentRepository.GetByName(ctx, name)
		if err != nil {
			return 0, err
		}
		return department.ID, nil
	}

	newDepartment := &entities.CreateDepartmentRequest{
		Name: name,
	}
	return s.DepartmentRepository.Create(ctx, newDepartment)
}
