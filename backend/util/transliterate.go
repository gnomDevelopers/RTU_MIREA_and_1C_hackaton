package util

import (
	"fmt"
	"github.com/mehanizm/iuliia-go"
	"strings"
)

func GenerateLogin(fio string) string {
	parts := strings.Fields(fio)
	if len(parts) < 2 {
		return "" // Если слишком короткое имя, то возвращаем пустую строку
	}

	// Применяем транслитерацию
	surname := transliterate(parts[0])
	name := string(transliterate(parts[1])[0])
	var thirdName string
	if len(parts) > 2 {
		thirdName = string(transliterate(parts[2])[0])
	}

	// Если отчество есть, то добавляем его в логин
	if thirdName != "" {
		return fmt.Sprintf("%s_%s_%s", strings.ToLower(surname), strings.ToLower(name), strings.ToLower(thirdName))
	}

	// Если отчества нет, то формируем логин без отчества
	return fmt.Sprintf("%s_%s", strings.ToLower(surname), strings.ToLower(name))
}

func transliterate(text string) string {
	transliteratedText := iuliia.Wikipedia.Translate(text)
	return transliteratedText
}
