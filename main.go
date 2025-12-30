package main

import (
	"net/http"
	simpleconnection "thirdApp/feature_postgres/simple_connection"
	"thirdApp/handlers"
)

func main() {
	simpleconnection.CheckConnection()
	
	http.HandleFunc("/add", handlers.AddTaskHandler)
	http.HandleFunc("/delete", handlers.DeleteTaskHandler)
	http.HandleFunc("/getTask", handlers.GetTaskByID)
	http.HandleFunc("/listTasks", handlers.GetAllTasksHandler)

	http.ListenAndServe(":9091", nil)

}
