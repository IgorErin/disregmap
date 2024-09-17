package disasterdb

import (
	"context"
	"log"

	"fetchdisasters/common"
	"fetchdisasters/jtable"

	"github.com/jmoiron/sqlx"
)

func QueryRegion(ctx context.Context, db *sqlx.DB) []common.Region {
	var res []common.Region

	err := db.Select(&res, "SELECT id, name from regions")
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func InsertCoordinates(coords []jtable.CoordRow, db *sqlx.DB) {
	command := `INSERT INTO coordinates 
		(id, longitude, latitude, region_id) 
		VALUES ($1, $2, $3, $4)
	`

	for _, c := range coords {
		db.MustExec(command, c.Id, c.Longitude, c.Latitude, c.RegionId)
	}
}
