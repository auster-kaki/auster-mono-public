package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type TravelSpotItinerary struct {
	db *bun.DB
}

func NewTravelSpotItinerary(db *bun.DB) *TravelSpotItinerary {
	return &TravelSpotItinerary{db: db}
}

func (t *TravelSpotItinerary) Create(ctx context.Context, ents ...entity.TravelSpotItinerary) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *TravelSpotItinerary) GetByTravelSpotID(ctx context.Context, travelSpotID entity.TravelSpotID) (entity.TravelSpotItineraries, error) {
	res := entity.TravelSpotItineraries{}
	if err := t.db.NewSelect().Model(&res).Where("travel_spot_id = ?", travelSpotID).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *TravelSpotItinerary) GetByTravelSpotIDs(ctx context.Context, travelSpotIDs []entity.TravelSpotID) (entity.TravelSpotItineraries, error) {
	res := entity.TravelSpotItineraries{}
	if err := t.db.NewSelect().Model(&res).Where("travel_spot_id IN (?)", bun.In(travelSpotIDs)).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *TravelSpotItinerary) GetByKinds(ctx context.Context, kinds []string) (entity.TravelSpotItineraries, error) {
	res := entity.TravelSpotItineraries{}
	if err := t.db.NewSelect().Model(&res).Where("kind IN (?)", bun.In(kinds)).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
