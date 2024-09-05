package disasterdb

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
)

type Disaster struct {
	Id             int    `db: id`
	RegionId       int    `db: region_id`
	TypeInfoId     int    `db: type_info_id`
	JScore         int    `db: j_score`
	Source         string `db: source`
	Description    string `db: description`
	DetailedSource string `db: detailed_source`

	YearStart  int `db: year_start`
	StartDay   int `db: start_day`
	StartMonth int `db: start_month`

	YearEnd  int `db: year_end`
	EndDay   int `db: end_day`
	EndMonth int `db: end_month`

	Casualties int `db: casualties`

	DamageAmount   int    `db: damage_amount`
	DamageCurrency string `db: damage_currency`
}

type TypeInfo struct {
	Id   int    `db: id`
	Name string `db: name`
}

type Region struct {
	Id   int    `db: id`
	Name string `db: name`
}

func QueryDisaster(ctx context.Context, db *sqlx.DB) []Disaster {
	rows, err := db.QueryContext(ctx, `SELECT * FROM disasters`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []Disaster

	for rows.Next() {
		err := sqlx.StructScan(rows, &result)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func QueryTypeInfo(ctx context.Context, db *sqlx.DB) []TypeInfo {
	return nil
}

func QueryRegion(ctx context.Context, db *sqlx.DB) []Region {
	return nil
}
