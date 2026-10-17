package simpleconnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

//"postgres://YourUserName:YourPassword@YourHostName:5432/YorDatabaseName"

func CheckConnection() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:3039306@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Connection to database success!")
}
