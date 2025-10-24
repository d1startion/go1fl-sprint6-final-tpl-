package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// min возвращает минимальное из двух чисел
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// IndexHandler — возвращает форму из index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Логируем запрос к главной странице
	fmt.Printf("Запрос к главной странице: %s %s\n", r.Method, r.URL.Path)
	http.ServeFile(w, r, "index.html")
}

// UploadHandler — обрабатывает загрузку файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Запрос к /upload: %s %s\n", r.Method, r.URL.Path)

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		fmt.Printf("Ошибка парсинга формы: %v\n", err)
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Printf("Ошибка получения файла: %v\n", err)
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fmt.Printf("Получен файл: %s, размер: %d байт\n", handler.Filename, handler.Size)

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Прочитано данных: %d байт\n", len(data))
	fmt.Printf("Содержимое файла (первые 100 символов): %q\n", string(data)[:min(100, len(string(data)))])

	result, err := service.AutoConvert(string(data))
	if err != nil {
		fmt.Printf("Ошибка конвертации: %v\n", err)
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Результат конвертации: %q\n", result)

	// Создаём локальный файл с результатом
	ext := filepath.Ext(handler.Filename)
	fileName := time.Now().UTC().Format("20060102_150405") + ext
	fmt.Printf("Попытка создания файла: %s\n", fileName)
	newFile, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Ошибка создания файла %s: %v\n", fileName, err)
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()
	fmt.Printf("Файл %s успешно создан\n", fileName)

	_, err = newFile.Write([]byte(result))
	if err != nil {
		fmt.Printf("Ошибка записи в файл %s: %v\n", fileName, err)
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Результат успешно записан в файл %s\n", fileName)

	// Возвращаем результат пользователю
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Результат конвертации:")
	fmt.Fprintln(w, result)
}
