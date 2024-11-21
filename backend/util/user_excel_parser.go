package util

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"server/internal/entities"
)

func ExcelFileParser(file multipart.File) (*[]entities.AddUserDataRequest, error) {
	// Открытие Excel файла
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть Excel файл: %w", err)
	}

	// Получаем первую страницу
	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать строки из Excel: %w", err)
	}

	// Массив для хранения данных
	var usersData []entities.AddUserDataRequest

	// Проходим по всем строкам Excel файла, начиная с 2-й
	for _, row := range rows[1:] {
		if len(row) < 7 {
			// Если строка неполная, пропускаем
			continue
		}

		userData := entities.AddUserDataRequest{
			LastName:   row[0], // Фамилия
			FirstName:  row[1], // Имя
			FatherName: row[2], // Отчество
			Group:      row[3], // Группа
			University: row[4], // Университет
			Faculty:    row[5], // Факультет
			Department: row[6], // Направление
		}

		usersData = append(usersData, userData)
	}

	return &usersData, nil
}
