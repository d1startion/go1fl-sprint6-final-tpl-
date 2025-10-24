package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "morse-converter: ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)

	logger.Println("Сервер запущен на http://localhost:8080")

	err := srv.Server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
