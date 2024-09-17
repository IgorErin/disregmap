package jtable

import (
	"fetchdisasters/dmap"
	"fetchdisasters/common"
)

type CoordRow struct {
	Id        int     `sql:"id"`
	Longitude float64 `sql:"longitude"`
	Latitude  float64 `sql:"latitude"`
	RegionId int	  `sql:"region_id"`
}

func Construct(regs []common.Region) []CoordRow {
	var coordinates []CoordRow 

	var coordCount int
	for _, r := range regs {
		coords := dmap.Lookup(r.Name)
		for _, c := range coords {
			coordinates = append(coordinates, CoordRow{Id: coordCount, Longitude: c.Longitude, Latitude: c.Latitude, RegionId: r.Id })
			coordCount++
		}
	}

	return coordinates
}
