package server

import (
	"go1fl-sprint6-final-tpl-mnn/internal/handlers"
	"log"
	"net/http"
	"time"
)

// 1. Создайте структуру сервера с полями для логгера (log.Logger) и http-сервера (http.Server).
type Server struct {
	Logger     *log.Logger
	HTTPServer http.Server
}

// 2. Создайте функцию, в которой нужно создать http-роутер. Функция принимает log.Logger и возвращает экземпляр структуры вашего сервера.
func NewServer(logger *log.Logger) *Server {
	// Создаем HTTP-роутер
	router := http.NewServeMux()

	// 3. Зарегистрируйте ваши хендлеры в http-роутере.
	router.HandleFunc("/upload", handlers.HandleUpload)
	router.HandleFunc("/", handlers.HandleMain)

	// Создайте экземпляр структуры http.Server. Для настройки вашего сервера используйте следующие поля:
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	//5. Верните ссылку на ваш сервер.
	return &Server{
		Logger:     logger,
		HTTPServer: *server,
	}
}
