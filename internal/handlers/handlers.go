package handlers

import (
	"fmt"
	"go1fl-sprint6-final-tpl-mnn/internal/service"
	"go1fl-sprint6-final-tpl-mnn/pkg/morse"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	var dst_data, dst_file string

	// 1. Парсить html-форму из файла index.html.
	r.ParseMultipartForm(10 << 20) // 10 MB

	// 2. Получить файл из формы (не забудьте его закрыть).
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	// закрываем файл
	defer file.Close()

	//3. Прочитать данные из файла.
	data, err := io.ReadAll(file)
	if err != nil {
		// Если произошла ошибка при чтении файла, возвращаем ошибку
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	//4. Передать эти данные в функцию автоопределения из пакета service, которую вы создали, чтобы получить переконвертируемую строку.
	strData := string(data)
	if service.IsMorseCode(strData) {
		dst_data = morse.ToText(strData)
	} else {
		dst_data = morse.ToMorse(strData)
	}

	// 5. Создать локальный файл. Эта операция обычно небезопасна и так делать не рекомендуется, но в рамках нашего задания хотелось бы более наглядного результата, поэтому мы решились
	// на этот шаг, ради видимого результата. А вообще, обычно используют временные файлы.
	// 6. Записать в локальный файл результат конвертации строки. Для генерации имени файла вы можете использовать время с помощью time.Now().UTC().String().
	//	Чтобы получить расширения файла, используйте filepath.Ext().
	ext := filepath.Ext(string(fileHeader.Filename))

	dst_file = time.Now().Format("2006-01-02T15:04:05Z") + ext

	if err := os.WriteFile(dst_file, []byte(dst_data), 0755); err != nil {
		http.Error(w, fmt.Sprintf("Error writing file: %v", err), http.StatusInternalServerError)
		return
	}

	// 7. Вернуть результат конвертации строки.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(dst_data))
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	// Читаем содержимое файла index.html
	data, err := os.ReadFile("index.html")
	if err != nil {
		// Если произошла ошибка при чтении файла, возвращаем ошибку
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}
