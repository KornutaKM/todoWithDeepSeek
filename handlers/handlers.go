package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"
)

var listTasks = make(map[int]Task)

// Используем заглавные буквы для экспорта полей JSON
type Task struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func (t *Task) validate() bool {
	return t.Text != ""
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Метод не поддерживается"))
		return
	}

	var task Task
	// Сначала декодируем тело запроса
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Ошибка чтения тела запроса: " + err.Error()
		w.Write([]byte(msg))
		return
	}
	defer r.Body.Close()

	// Валидируем задачу
	if !task.validate() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Текст задачи не может быть пустым"))
		return
	}

	// Генерируем уникальный ID
	id := rand.IntN(10000)
	for {
		if _, exists := listTasks[id]; exists {
			id = rand.IntN(10000)
		} else {
			break
		}
	}

	// Устанавливаем ID и время создания
	task.ID = id
	task.CreatedAt = time.Now()

	// Сохраняем задачу
	listTasks[id] = task

	// Отправляем успешный ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	// Возвращаем созданную задачу
	response, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка при формировании ответа"))
		return
	}
	w.Write(response)

	// Логирование (для отладки)
	fmt.Printf("Добавлена задача ID %d. Всего задач: %d\n", id, len(listTasks))
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Лучше использовать параметры URL или заголовки для передачи ID
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Метод не поддерживается"))
		return
	}

	// Получаем ID из query параметра
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		// Альтернативно: из тела запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Ошибка чтения тела запроса"))
			return
		}
		defer r.Body.Close()
		idStr = string(body)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Некорректный ID"))
		return
	}

	// Проверяем существование задачи
	if _, ok := listTasks[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Задача не найдена"))
		return
	}

	// Удаляем задачу
	delete(listTasks, id)

	// Отправляем ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Задача удалена"))

	// Логирование
	fmt.Printf("Удалена задача ID %d. Осталось задач: %d\n", id, len(listTasks))
}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Метод не поддерживается"))
		return
	}

	// Подготавливаем список задач
	tasksList := make([]Task, 0, len(listTasks))
	for _, task := range listTasks {
		tasksList = append(tasksList, task)
	}

	// Сериализуем в JSON
	w.Header().Set("Content-Type", "application/json")
	if len(tasksList) == 0 {
		w.Write([]byte("[]"))
		return
	}

	jsonData, err := json.Marshal(tasksList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка при формировании ответа"))
		return
	}

	w.Write(jsonData)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Метод не поддерживается"))
		return
	}

	// Получаем ID из query параметра
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID не указан"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Некорректный ID"))
		return
	}

	// Ищем задачу
	task, ok := listTasks[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Задача не найдена"))
		return
	}

	// Возвращаем задачу в формате JSON
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка при формировании ответа"))
		return
	}

	w.Write(jsonData)
}

// Дополнительно: функция для обновления задачи
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Метод не поддерживается"))
		return
	}

	// Получаем ID из query параметра
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID не указан"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Некорректный ID"))
		return
	}

	// Проверяем существование задачи
	if _, ok := listTasks[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Задача не найдена"))
		return
	}

	// Декодируем обновленные данные
	var updatedTask Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка чтения тела запроса"))
		return
	}
	defer r.Body.Close()

	// Валидируем
	if !updatedTask.validate() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Текст задачи не может быть пустым"))
		return
	}

	// Обновляем задачу (сохраняем ID и время создания)
	currentTask := listTasks[id]
	currentTask.Text = updatedTask.Text
	listTasks[id] = currentTask

	// Отправляем обновленную задачу
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(currentTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка при формировании ответа"))
		return
	}

	w.Write(jsonData)
}