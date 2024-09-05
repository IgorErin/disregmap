package disasterdb

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type Disaster struct {
	Id       int `db:"id"`
	RegionId int `db:"region_id"`
	// TODO(add fields)
	// TypeInfoId     int    `db: type_info_id`
	// JScore         int    `db: j_score`
	// Source         string `db: source`
	// Description    string `db: description`
	// DetailedSource string `db: detailed_source`

	YearStart  int `db:"year_start"`
	StartDay   int `db:"start_day"`
	StartMonth int `db:"start_month"`

	YearEnd  int `db:"year_end"`
	EndDay   int `db:"end_day"`
	EndMonth int `db:"end_month"`

	// Casualties int `db: casualties`

	// DamageAmount   int    `db: damage_amount`
	// DamageCurrency string `db: damage_currency`
}

func QueryDisaster(ctx context.Context, db *sqlx.DB) []Disaster {
	query := `
	SELECT
		id, region_id, year_start,
		start_day, start_month, year_end,
		end_day, end_month
	FROM disasters`

	var res []Disaster

	err := db.Select(&res, query)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

type TypeInfo struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func QueryTypeInfo(ctx context.Context, db *sqlx.DB) []TypeInfo {
	var res []TypeInfo

	err := db.Select(&res, "SELECT id, name from type_infos")
	if err != nil {
		log.Fatal(err)
	}

	return res
}

type Region struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func QueryRegion(ctx context.Context, db *sqlx.DB) []Region {
	var res []Region

	err := db.Select(&res, "SELECT id, name from type_infos")
	if err != nil {
		log.Fatal(err)
	}

	return res
}
