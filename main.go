package main

import (
	"context"
	"fmt"
	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 5 {
			task.Title = "Pokormit koshku"
			task.Description = "Otsypat koshke 30 gramm korma"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}

			break
		}
	}

	fmt.Println("succeed!")
}
