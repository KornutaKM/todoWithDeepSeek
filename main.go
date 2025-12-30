package main

import (
	"net/http"
	"thirdApp/handlers"
)

func main() {
	http.HandleFunc("/add", handlers.AddTaskHandler)
	http.HandleFunc("/delete", handlers.DeleteTaskHandler)
	http.HandleFunc("/getTask", handlers.GetTaskByID)
	http.HandleFunc("/listTasks", handlers.GetAllTasksHandler)

	http.ListenAndServe(":9091", nil)
}