package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type TravelSpotItineraryItem struct {
	db *bun.DB
}

func NewTravelSpotItineraryItem(db *bun.DB) *TravelSpotItineraryItem {
	return &TravelSpotItineraryItem{db: db}
}

func (t *TravelSpotItineraryItem) Create(ctx context.Context, ents ...entity.TravelSpotItineraryItem) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *TravelSpotItineraryItem) GetByTravelSpotItineraryIDs(ctx context.Context, travelSpotItineraryIDs []entity.TravelSpotItineraryID) (entity.TravelSpotItineraryItems, error) {
	res := entity.TravelSpotItineraryItems{}
	if err := t.db.NewSelect().Model(&res).Where("travel_spot_itinerary_id IN (?)", bun.In(travelSpotItineraryIDs)).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
