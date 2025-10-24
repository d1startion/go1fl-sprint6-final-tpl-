package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// NewServer — создаёт новый сервер и регистрирует маршруты
func NewServer(logger *log.Logger) *Server {
	fmt.Println("Регистрация маршрутов...")
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	fmt.Println("Зарегистрирован маршрут: /")
	mux.HandleFunc("/upload", handlers.UploadHandler)
	fmt.Println("Зарегистрирован маршрут: /upload")

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: s,
	}
}
