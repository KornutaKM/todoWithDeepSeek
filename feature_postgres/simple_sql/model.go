package simple_sql

import "time"

type TaskModel struct {
	Id        int
	Text      string
	CreatedAt time.Time
}
