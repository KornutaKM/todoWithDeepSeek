package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		text VARCHAR(1000) NOT NULL,
		created_at TIMESTAMP NOT NULL,

		UNIQUE(text)
	);
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
