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

// IndexHandler — возвращает форму из index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler — обрабатывает загрузку файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	result, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаём локальный файл с результатом
	ext := filepath.Ext(handler.Filename)
	fileName := time.Now().UTC().Format("20060102_150405") + ext
	newFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.Write([]byte(result))
	if err != nil {
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат пользователю
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Результат конвертации:")
	fmt.Fprintln(w, result)
}
