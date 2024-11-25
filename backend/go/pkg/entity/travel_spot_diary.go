package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type TravelSpotDiaryID string

type TravelSpotDiary struct {
	ID           TravelSpotDiaryID
	UserID       UserID
	TravelSpotID TravelSpotID
	Title        string
	Date         time.Time
	PhotoPath    string
	Description  string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:travel_spot_diary,alias:tsd"`
}

type TravelSpotDiaries []*TravelSpotDiary
