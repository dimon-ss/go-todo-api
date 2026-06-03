package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
	"github.com/dimon-ss/go-todo-api/internal/handler"
	"github.com/dimon-ss/go-todo-api/internal/repository"
)

func main() {
	db, err := sql.Open("sqlite", "./blog.db")
	if err != nil {
		log.Fatal("Ошибка открытия БД:", err)
	}
	defer db.Close()

	// Создание таблицы
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	h := handler.NewHandler(repo)

	fmt.Println("✅ ToDo REST API запущен на http://localhost:8080")
	fmt.Println("GET    /tasks     → получить все задачи")
	fmt.Println("POST   /tasks     → создать задачу")
	fmt.Println("DELETE /tasks/1   → удалить задачу")

	http.HandleFunc("/tasks", h.TasksHandler)
	http.HandleFunc("/tasks/", h.TaskHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
