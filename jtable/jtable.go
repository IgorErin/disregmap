package jtable

import (
	"context"
	"db/common"
	"db/disasterdb"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

func timeFormat(date disasterdb.Date) string {
	return fmt.Sprintf("%d/%d/%d", date.Day, date.Moth, date.Year)
}

func lookupCoordinates(_ string) string {
	fst := rand.Float64() * 60
	snd := rand.Float64() * 60

	return fmt.Sprintf("%f, %f", fst, snd)
}

type row struct {
	Titel string `json:"Titel"`
	Start string `json:"Start"`
	End string `json:"End"`
	Description string `json:"Description"`
	Image string `json:"Image"`
	Place string `json:"Place"`
	Location string `json:"Location"`
	Source string `json:"Source"`
	SourceURL string `json:"Source URL"`
}

func toRow(act disasterdb.Action) row {
	startDate := timeFormat(act.StartDate) 
	return row {
		Titel: "Disaster",
		Start: startDate,
		End: startDate, // for now, strange values in tables
		Description: act.Description,
		Image: "",
		Place: "",
		Location: lookupCoordinates(act.Region),
		Source: "",
		SourceURL: "",
	}
}

func ToJson(_ context.Context, acts []disasterdb.Action) []byte {
	rsl := common.Fmap(func (act disasterdb.Action ) row { return toRow(act ) }, acts)

	res, err := json.Marshal(rsl)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
