package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Сделать запрос
	response, err := http.Get("https://go.dev")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Создайте файл вывода
	outFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Копирование данных из ответа HTTP в файл
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
