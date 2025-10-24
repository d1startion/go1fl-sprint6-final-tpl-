package main

import (
	"fmt"
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	fmt.Println("Запуск сервера...")
	logger := log.New(log.Writer(), "morse-converter: ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)
	fmt.Println("Сервер создан, запуск...")

	logger.Println("Сервер запущен на http://localhost:8080")

	err := srv.Server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
