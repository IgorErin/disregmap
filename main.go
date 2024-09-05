// main.go
package main

import (
	"context"
	"database/sql"
	"log"

	"encoding/json"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func query(ctx context.Context, db *sql.DB) {
//     rows, err := db.QueryContext(ctx, `SELECT table_name
//   FROM information_schema.tables
//  WHERE table_schema='public'
//    AND table_type='BASE TABLE';`)

    rows, err := db.QueryContext(ctx, `SELECT * FROM disasters`)

    if err != nil {
        log.Fatal(err.Error())
    }    
    defer rows.Close()

    cols, err := rows.Columns()
    if err != nil {
        log.Fatal(err.Error())
    }

    allgeneric := make([]map[string]interface{}, 0)
    colvals := make([]interface{}, len(cols))

    for rows.Next() {
        colassoc := make(map[string]interface{}, len(cols))

        for i := range colvals {
            colvals[i] = new(interface{})
        }
        if err := rows.Scan(colvals...); err != nil {
            log.Fatal(err.Error())
        }
        for i, col := range cols {
            colassoc[col] = *colvals[i].(*interface{})
        }
        allgeneric = append(allgeneric, colassoc)
    }

    rows.Close()

    j, err := json.Marshal(allgeneric)
    if err != nil{
        panic(err)
    }
    fmt.Println(string(j))
}

func main() {
    db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err.Error())
    }
    defer db.Close()

    ctx := context.Background()

    if err := db.PingContext(ctx); err != nil {
        log.Fatal(err.Error())
    }

    query(ctx, db)
}

