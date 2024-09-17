// main.go
package main

import (
	"context"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"fetchdisasters/disasterdb"
	"fetchdisasters/jtable"
)

func main() {
	connectstr := os.Getenv("DATABASE_URL")
	if connectstr == "" {
		log.Fatal("No database url specified!")
	}

	db, err := sqlx.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	ctx := context.Background()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err.Error())
	}

	regs := disasterdb.QueryRegion(ctx, db)
	coordRows := jtable.Construct(regs)
	disasterdb.InsertCoordinates(coordRows, db)
}
