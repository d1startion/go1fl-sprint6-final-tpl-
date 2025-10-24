package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// AutoConvert определяет тип текста и выполняет конвертацию
func AutoConvert(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("пустой ввод")
	}

	// Если строка содержит только точки, тире и пробелы — это Морзе
	if strings.ContainsAny(input, ".-") && !strings.ContainsAny(input, "абвгдАБВГДabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return morse.ToText(input), nil
	}

	// Если содержит буквы — значит обычный текст
	if strings.ContainsAny(input, "абвгдАБВГДabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return morse.ToMorse(input), nil
	}

	return "", errors.New("не удалось определить тип данных")
}
