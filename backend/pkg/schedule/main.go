package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"server/internal/entities"
	"strconv"
	"strings"
)

func main() {
	Parse()
}

type Class struct {
	Id          int      `json:"class_id"`                 // Id пары
	Groups      []int    `json:"class_group_ids"`          // Список Id групп, для которых пара
	GroupsNames []string `json:"class_group_names"`        // список названий групп, т.к. не все группы есть в таблице учебных групп
	Teacher     int      `json:"class_teacher_id"`         // Id преподавателя, который ведет пару
	TeacherName string   `json:"class_teacher_name"`       // фио преподавателя
	Count       int      `json:"class_count"`              // Какая пара по счету за день
	Weekday     int      `json:"class_weekday"`            // Номер дня недели
	Week        int      `json:"class_week"`               // Номер учебной неделяя
	Name        string   `json:"class_name"`               // Название пары
	Type        string   `json:"class_type"`               // Тип пары
	Location    string   `json:"class_location,omitempty"` // Местонахождение
}

type StudentGroup struct {
	Id        int    // Уникальный идентификатор группы
	Grade     int    // Курс группы
	Institute string // Название института
	Name      string // Название группы
	Students  []int  // Идентификаторы студентов в группе
}

func Parse() error {
	path := "pkg/schedule/excel_files" //Добавление в директорию нужного института
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	//Проход по файлам экселя
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xlsx") {
			filePath := filepath.Join(path, file.Name()) //Получение имени файла

			xlFile, err := excelize.OpenFile(filePath) //Открытие файла
			if err != nil {
				fmt.Println("Ошибка открытия файла:", err)
				continue
			}

			rows, _ := xlFile.GetRows("Расписание") //Получение листа Excel
			find2(rows)                             //Получение нового ID и добавление одного файла в БД
		}
	}

	return nil
}

func find2(rows [][]string) {
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
	}
}
