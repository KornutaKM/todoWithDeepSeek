package simpleconnection

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() error {
	// Используем контекст с таймаутом для предотвращения зависания
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Подключение к базе данных
	connString := "postgres://koxypo:admin@192.168.1.224:5432/postgres"
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к БД: %w", err)
	}
	defer conn.Close(ctx)

	// Проверка соединения
	if err := conn.Ping(ctx); err != nil {
		return fmt.Errorf("не удалось проверить соединение с БД: %w", err)
	}

	fmt.Println("Подключение к базе данных успешно!")
	return nil
}
