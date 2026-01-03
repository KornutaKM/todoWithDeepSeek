package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
		SELECT id, text, created_at
		FROM tasks
		ORDER BY id ASC
		`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		err := rows.Scan(&task.Id, &task.Text, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		// printTask(task)
	}

	return tasks, nil
}

func printTask(task TaskModel) {
	fmt.Println("---------------------------------------------")
	fmt.Println("id: ", task.Id)
	fmt.Println("text: ", task.Text)
	fmt.Println("createdAt: ", task.CreatedAt)
}
