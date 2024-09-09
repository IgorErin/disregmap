// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"db/disasterdb"
	"db/jtable"
	// "strings"
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

	disasters := disasterdb.Retriev(ctx, db)[:100]
	jsonbytes := jtable.ToJson(ctx, disasters)

	fmt.Println(string(jsonbytes))
}
