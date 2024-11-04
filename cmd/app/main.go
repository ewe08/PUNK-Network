package main

import (
	"PUNK-Network/pkg/repository"
	"log"
	"net/http"
)

func main() {
	// Инициализация базы данных
	db := repository.GetDB()
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
