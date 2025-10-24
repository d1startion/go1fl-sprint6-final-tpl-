package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// AutoConvert определяет тип текста и выполняет конвертацию
func AutoConvert(input string) (string, error) {
	fmt.Printf("AutoConvert вызван с входными данными: %q\n", input)
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Printf("Пустой ввод\n")
		return "", errors.New("пустой ввод")
	}

	// Если строка содержит только точки, тире и пробелы — это Морзе
	// Проверяем, содержит ли строка только символы Морзе
	isMorse := true
	for _, r := range input {
		if r != '.' && r != '-' && r != ' ' {
			isMorse = false
			break
		}
	}

	// Проверяем, содержит ли строка буквы
	containsLetters := false
	for _, r := range input {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') {
			containsLetters = true
			break
		}
	}

	fmt.Printf("Проверка на Морзе: состоит только из символов Морзе? %v, содержит буквы? %v\n", isMorse, containsLetters)

	if isMorse && !containsLetters {
		fmt.Printf("Вызов morse.ToText для: %q\n", input)
		result := morse.ToText(input)
		fmt.Printf("Результат morse.ToText: %q\n", result)
		return result, nil
	}

	// Если содержит буквы — значит обычный текст
	fmt.Printf("Проверка на обычный текст: содержит буквы? %v\n", containsLetters)

	if containsLetters {
		fmt.Printf("Вызов morse.ToMorse для: %q\n", input)
		result := morse.ToMorse(input)
		fmt.Printf("Результат morse.ToMorse: %q\n", result)
		return result, nil
	}

	fmt.Printf("Не удалось определить тип данных для: %q\n", input)
	return "", errors.New("не удалось определить тип данных")
}
