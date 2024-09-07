package jtable

import (
	"context"
	"db/disasterdb"
	"encoding/json"
	"fmt"
	"log"

	// "cloud.google.com/go/translate"
	// "golang.org/x/text/language"
)

// TODO
// 1 translate region
// 2 convert time

// func transReg(ctx context.Context, reg string) string {
// 	client, err := translate.NewClient(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	translations, err := client.Translate(ctx,
// 		[]string{reg}, language.English,
// 		&translate.Options{
// 			Source: language.Russian,
// 			Format: translate.Text,
// 		})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return translations[0].Text
// } 

func timeFormat(date disasterdb.Date) string {
	return fmt.Sprintf("%d/%d/%d", date.Day, date.Moth, date.Year)
}

type row struct {
	Titel string `json:"Titel"`
	Start string `json:"Start"`
	End string `json:"End"`
	Description string `json:"Description"`
	Image string `json:"Image"`
	Place string `json:"Place"`
	Location string `json:"Location_Original"`
}

func toRow(ctx context.Context, act disasterdb.Action) row {
	return row {
		Titel: "",
		Start: timeFormat(act.StartDate),
		End: timeFormat(act.EndDate),
		Description: "",
		Image: "",
		Place: act.Region, // transReg(ctx, act.Region),
		Location: "",
	}
}

func fmap[A any, B any](f func (A) B, sl []A) []B {
	res := make([]B, len(sl))
	for ind, item := range sl {
		res[ind] = f(item)
	}
	return res
}

func ToJson(ctx context.Context, acts []disasterdb.Action) []byte {
	rsl := fmap(func (act disasterdb.Action ) row { return toRow(ctx, act ) }, acts)

	res, err := json.Marshal(rsl)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
