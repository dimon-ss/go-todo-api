package repository

import (
	"database/sql"
	"github.com/dimon-ss/go-todo-api/internal/model"
)

type Repository interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task *model.Task) error
	DeleteTask(id int) error
}

type SQLiteRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) GetAllTasks() ([]model.Task, error) {
	rows, err := r.db.Query("SELECT id, title, content, created_at FROM tasks ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *SQLiteRepository) CreateTask(task *model.Task) error {
	_, err := r.db.Exec(
		"INSERT INTO tasks (title, content, created_at) VALUES (?, ?, ?)",
		task.Title, task.Content, task.CreatedAt,
	)
	return err
}

func (r *SQLiteRepository) DeleteTask(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
