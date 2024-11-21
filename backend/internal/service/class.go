package service

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"server/internal/entities"
	"server/internal/repository"
	"strconv"
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

func (s *ClassService) GetByTeacherName(c context.Context, teacherName string) (*[]entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.repository.GetByTeacherName(ctx, teacherName)
	return classes, err
}

func (s *ClassService) GetByAuditoryId(c context.Context, auditoryId int) (*[]entities.Class, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	classes, err := s.repository.GetByAuditoryId(ctx, auditoryId)
	return classes, err
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

func (s *ClassService) Parse(fileName string) error {
	path := "./backend/tmp"                   //Добавление в директорию нужного института
	filePath := filepath.Join(path, fileName) //Получение имени файла

	xlFile, err := excelize.OpenFile(filePath) //Открытие файла
	if err != nil {
		return err
	}

	rows, _ := xlFile.GetRows("Расписание") //Получение листа Excel
	err = s.find(rows)                      //Получение нового ID и добавление одного файла в БД
	if err != nil {
		return err
	}

	return nil
}

func (s *ClassService) find(rows [][]string) error {
	countDay, last := 0, "-1"
	c := 0
	for j := 5; j < len(rows[1]); j += 6 {
		group := rows[0][j]
		var classes []entities.Class
		for i := 2; i < len(rows); i++ {
			var class entities.Class
			if len(rows[i]) <= j+5 || rows[i][j] == "" {
				continue
			}
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
			class.Date = rows[i][j+5]
			class.Weekday = countDay + 1
			class.Week, _ = strconv.Atoi(rows[i][1])
			if last > rows[i][1] {
				countDay++
			}
			class.TimeStart = rows[i][2]
			class.TimeEnd = rows[i][3]

			classes = append(classes, class)
			last = rows[i][1]
		}
		c++
		last = "-1"
		fmt.Println(classes)
		if c == 2 {
			fmt.Println()
		}

		fmt.Println("-------------------------------------------")
		_, err := s.repository.Create(context.Background(), &classes)
		if err != nil {
			return err
		}
		fmt.Println("Ok")
	}

	return nil
}
