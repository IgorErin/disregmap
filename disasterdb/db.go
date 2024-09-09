package disasterdb

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type Disaster struct {
	Id       int `db:"id"`
	RegionId int `db:"region_id"`
	TypeInfoId     int    `db:"type_info_id"`
	// JScore         float32    `db:"j_score"`
	// Source         string `db:"source"`
	Description    string `db:"description"`
	// DetailedSource string `db:"detailed_source"`

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
		end_day, end_month,
		description
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

	err := db.Select(&res, "SELECT id, name from regions")
	if err != nil {
		log.Fatal(err)
	}

	return res
}

type Date struct {
	Day int 
	Moth int 
	Year int
}

type Action struct {
	Region string
	TypeInfo string

	StartDate Date  
	EndDate Date

	Description string
	JScore float32
}

func Retriev(ctx context.Context, db *sqlx.DB) []Action {
	disasters := QueryDisaster(ctx, db)

	typeInfo := QueryTypeInfo(ctx, db)
	typeInfoMap := func () map[int]string {
		m := make(map[int]string, len(typeInfo))
		for _, info := range typeInfo {
			m[info.Id] = info.Name
		}
		return m
	}()

	region := QueryRegion(ctx, db)
	regionMap := func () map[int]string {
		m := make(map[int]string, len(region))
		for _, r := range region {
			m[r.Id] = r.Name
		}
		return m
	}()

	results := make([]Action, 0, len(disasters))
	for _, d := range disasters {
		results = append(results, Action{ 
			Region: regionMap[d.RegionId],
			TypeInfo: typeInfoMap[d.TypeInfoId],

			StartDate: Date {Day: d.StartDay, Moth: d.StartMonth, Year: d.YearStart },
			EndDate: Date {Day: d.EndDay, Moth: d.EndMonth, Year: d.YearEnd },

			Description: d.Description,
		})
	}

	return results
}
