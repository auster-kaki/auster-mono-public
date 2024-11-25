package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type TravelSpot struct {
	db *bun.DB
}

func NewTravelSpot(db *bun.DB) *TravelSpot {
	return &TravelSpot{db: db}
}

func (t *TravelSpot) Create(ctx context.Context, ents ...entity.TravelSpot) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *TravelSpot) GetByVendorID(ctx context.Context, vendorID entity.VendorID) (entity.TravelSpots, error) {
	res := entity.TravelSpots{}
	if err := t.db.NewSelect().Model(&res).Where("vendor_id = ?", vendorID).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *TravelSpot) FindByID(ctx context.Context, id entity.TravelSpotID) (*entity.TravelSpot, error) {
	res := &entity.TravelSpot{}
	if err := t.db.NewSelect().Model(res).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *TravelSpot) GetByIDs(ctx context.Context, ids []entity.TravelSpotID) (entity.TravelSpots, error) {
	res := entity.TravelSpots{}
	if err := t.db.NewSelect().Model(&res).Where("id IN (?)", bun.In(ids)).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
