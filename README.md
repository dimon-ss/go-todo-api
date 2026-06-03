# Go Todo API

Простой REST API для управления задачами, написанный на Go.

## Возможности
- Получить список всех задач (GET /tasks)
- Создать новую задачу (POST /tasks)
- Удалить задачу (DELETE /tasks/{id})

## Как запустить

```bash
go run main.go
cat >> README.md << 'EOF'

## Примеры запросов

**Получить все задачи:**
```bash
curl http://localhost:8080/tasks
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить продукты","content":"Молоко, хлеб, яйца"}'
curl -X DELETE http://localhost:8080/tasks/1
curl -X DELETE http://localhost:8080/tasks/1

---

### 2. Создай `internal/model/task.go`

```bash
cat > internal/model/task.go << 'EOF'
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
