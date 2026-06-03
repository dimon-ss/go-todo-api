# Go Todo REST API

Простой REST API для управления задачами, написанный на Go с использованием SQLite.

## Возможности
- Получение списка всех задач (`GET /tasks`)
- Создание новой задачи (`POST /tasks`)
- Удаление задачи (`DELETE /tasks/{id}`)

## Как запустить

```bash
go run main.go
Сервер запускается на http://localhost:8080
Примеры запросов
1. Получить все задачи
Bashcurl http://localhost:8080/tasks
2. Создать новую задачу
Bashcurl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Купить продукты",
    "content": "Молоко, хлеб, яйца и сыр"
  }'
3. Удалить задачу (например, с id = 1)
Bashcurl -X DELETE http://localhost:8080/tasks/1
Стек технологий

Go + net/http
SQLite (modernc.org/sqlite)
JSON

Структура проекта

main.go — главный файл
internal/handler/ — обработчики запросов
internal/repository/ — работа с базой данных
internal/model/ — модели данных
