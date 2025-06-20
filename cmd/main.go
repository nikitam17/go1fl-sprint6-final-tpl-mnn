package main

import (
	"log"
	"os"

	"go1fl-sprint6-final-tpl-mnn/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Создаем сервер с помощью функции NewServer
	srv := server.NewServer(logger)

	// Запускаем сервер
	err := srv.HTTPServer.ListenAndServe()
	if err != nil {
		// Если при запуске сервера возникают ошибки, выводим их с помощью логгера на уровне Fatal
		logger.Fatal("Ошибка при запуске сервера: ", err)
	}
}
