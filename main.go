// main.go
package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"os"

	"db/disasterdb"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sqlx.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	ctx := context.Background()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err.Error())
	}

	res := disasterdb.QueryDisaster(ctx, db)

	fmt.Println(res)
}
