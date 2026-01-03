package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	// Используем контекст с таймаутом для предотвращения зависания
	

	// Подключение к базе данных
	connString := "postgres://koxypo:admin@192.168.1.224:5432/postgres"
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к БД: %w", err)
	}

	// Проверка соединения
	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("не удалось проверить соединение с БД: %w", err)
	}

	fmt.Println("Подключение к базе данных успешно!")
	return conn, nil
}
