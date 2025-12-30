package simpleconnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	connection, err := pgx.Connect(ctx, "postgres://koxypo:admin@192.168.1.224:5432/mydb")
	if err != nil {
		fmt.Println("Ошибка при подключении к БД:", err)
		panic(err)
	}

	if err := connection.Ping(ctx); err != nil {
		fmt.Println("Ошибка при подключении к БД:", err)
		panic(err)
	}

	fmt.Println("Подключение к базе данных успешно!")
}
