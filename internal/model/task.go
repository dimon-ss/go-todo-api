package model

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CreateTaskRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
