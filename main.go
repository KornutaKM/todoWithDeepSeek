package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	simple_connection "thirdApp/feature_postgres/simple_connection"
	"thirdApp/feature_postgres/simple_sql"
	"thirdApp/handlers"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// if err := simple_sql.InsertRow(ctx, conn, "Ужин2", time.Now()); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.UpdateRow(ctx, conn, 3, "hello"); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.DeleteRow(ctx, conn, 3); err != nil {
	// 	panic(err)
	// }

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	pp.Println(tasks)

	fmt.Println("Успех!")

	http.HandleFunc("/add", handlers.AddTaskHandler)
	http.HandleFunc("/delete", handlers.DeleteTaskHandler)
	http.HandleFunc("/getTask", handlers.GetTaskByID)
	http.HandleFunc("/listTasks", handlers.GetAllTasksHandler)

	http.ListenAndServe(":9091", nil)

}
