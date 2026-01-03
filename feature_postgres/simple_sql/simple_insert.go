package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlQuery := `
	INSERT INTO tasks (text, created_at)
	VALUES ($1, $2);
`

	_, err := conn.Exec(ctx, sqlQuery, task.Text, task.CreatedAt)

	return err
}
