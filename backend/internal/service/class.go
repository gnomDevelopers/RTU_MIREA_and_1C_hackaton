package service

import (
	"context"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"server/internal/entities"
	"server/internal/repository"
	"strings"
	"time"
)

type ClassService struct {
	repository repository.ClassRepository
	timeout    time.Duration
}

func NewClassService(repository repository.ClassRepository) *ClassService {
	return &ClassService{
		repository: repository,
		timeout:    time.Duration(10) * time.Second,
	}
}

func (s *ClassService) Create(c context.Context, class *[]entities.Class) ([]int, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	ids, err := s.repository.Create(ctx, class)
	return ids, err
}

func (s *ClassService) GetById(c context.Context, id int) (*entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	class, err := s.repository.GetById(ctx, id)
	return class, err
}

func (s *ClassService) GetByGroupName(c context.Context, groupName string) (*[]entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.repository.GetByGroupName(ctx, groupName)
	return classes, err
}

func (s *ClassService) GetByTeacherName(c context.Context, teacherName, university string) (*[]entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.repository.GetByTeacherName(ctx, teacherName, university)
	return classes, err
}

func (s *ClassService) GetByName(c context.Context, name, university string) (*[]entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.repository.GetByName(ctx, name, university)
	return classes, err
}

func (s *ClassService) GetByAuditory(c context.Context, auditory string) (*[]entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.repository.GetByAuditory(ctx, auditory)
	return classes, err
}

func (s *ClassService) SearchGroups(c context.Context, university string) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	groups, err := s.repository.SearchGroups(ctx, university)
	return groups, err
}

func (s *ClassService) SearchTeachers(c context.Context, university string) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	groups, err := s.repository.SearchTeachers(ctx, university)
	return groups, err
}

func (s *ClassService) SearchNames(c context.Context, university string) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	groups, err := s.repository.SearchNames(ctx, university)
	return groups, err
}

func (s *ClassService) Update(c context.Context, class *entities.Class) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Update(ctx, class)
	return err
}

func (s *ClassService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.Delete(ctx, id)
	return err
}

func (s *ClassService) Parse(fileName, university string) error {
	path := "./backend/tmp"
	filePath := filepath.Join(path, fileName)

	xlFile, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}

	rows, _ := xlFile.GetRows("Расписание")
	err = s.find(rows, university)
	if err != nil {
		return err
	}

	return nil
}

func (s *ClassService) find(rows [][]string, university string) error {
	countDay, last := 0, "-1"
	for j := 5; j < len(rows[1]); j += 6 {
		group := rows[0][j]
		var classes []entities.Class
		for i := 2; i < len(rows); i++ {
			if len(rows[i]) <= j+5 || rows[i][j] == "" {
				continue
			}
			dateStr := rows[i][j+5]
			layout := "01-02-06"
			startDate, err := time.Parse(layout, dateStr)
			if err != nil {
				return err
			}
			endDateStr := "12-23-24"
			endDate, _ := time.Parse(layout, endDateStr)
			for {
				var class entities.Class
				class.Name = rows[i][j]
				groupNames := rows[i][j+4]
				if groupNames == "" {
					class.GroupNames = append(class.GroupNames, group)
				} else {
					groupNamesSlice := strings.Split(groupNames, ",")
					for i := range groupNamesSlice {
						class.GroupNames = append(class.GroupNames, strings.Trim(groupNamesSlice[i], " "))
					}
				}
				class.TeacherNames = append(class.TeacherNames, rows[i][j+2])
				class.Type = rows[i][j+1]
				class.Date = dateStr
				class.Auditory = rows[i][j+3]
				if last > rows[i][1] {
					countDay++
				}
				class.Weekday = countDay + 1
				if rows[i][4] == "I" {
					class.Week = 1
				} else {
					class.Week = 2
				}
				class.TimeStart = rows[i][2]
				class.TimeEnd = rows[i][3]
				class.UniversityStr = university
				classes = append(classes, class)

				newDate := startDate.AddDate(0, 0, 14)

				if newDate.After(endDate) {
					break
				}

				startDate = newDate
				dateStr = newDate.Format(layout)
			}
			last = rows[i][1]
		}
		last = "-1"
		countDay = 0
		_, err := s.repository.Create(context.Background(), &classes)
		if err != nil {
			return err
		}
	}

	return nil
}
