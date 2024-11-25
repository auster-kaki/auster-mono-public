package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type ReservationID string

type Reservation struct {
	ID                ReservationID
	UserID            UserID
	TravelSpotID      TravelSpotID
	TravelSpotDiaryID TravelSpotDiaryID
	FromDate          time.Time
	ToDate            time.Time
	IsOffer           bool

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:reservation,alias:r"`
}

type Reservations []*Reservation

func (e *Reservations) TravelSpotIDs() []TravelSpotID {
	ids := make([]TravelSpotID, 0, len(*e))
	for _, r := range *e {
		ids = append(ids, r.TravelSpotID)
	}
	return ids
}

func (e *Reservations) TravelSpotDiaryIDs() []TravelSpotDiaryID {
	ids := make([]TravelSpotDiaryID, 0, len(*e))
	for _, r := range *e {
		ids = append(ids, r.TravelSpotDiaryID)
	}
	return ids
}
