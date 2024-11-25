package entity

import (
	"github.com/uptrace/bun"
)

type TravelSpotItineraryItemID string

type TravelSpotItineraryItem struct {
	ID                    TravelSpotItineraryItemID
	TravelSpotItineraryID TravelSpotItineraryID
	Name                  string
	Number                int

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:travel_spot_itinerary_item,alias:tsii"`
}

type TravelSpotItineraryItems []*TravelSpotItineraryItem
