package main

import (
	"log"
	"net/http"
	simpleconnection "thirdApp/feature_postgres/simple_connection"
	"thirdApp/handlers"
)

func main() {
	if err := simpleconnection.CheckConnection(); err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}

	http.HandleFunc("/add", handlers.AddTaskHandler)
	http.HandleFunc("/delete", handlers.DeleteTaskHandler)
	http.HandleFunc("/getTask", handlers.GetTaskByID)
	http.HandleFunc("/listTasks", handlers.GetAllTasksHandler)

	http.ListenAndServe(":9091", nil)

}
