package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, id int, newText string) error {

	sqlQuery := `
	UPDATE tasks
	SET text = $2
	WHERE id = $1;
	`

	_, err := conn.Exec(ctx, sqlQuery, id, newText)

	return err
}
