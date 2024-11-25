package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type Reservation struct {
	db *bun.DB
}

func NewReservation(db *bun.DB) *Reservation { return &Reservation{db: db} }

func (t *Reservation) Create(ctx context.Context, ents ...*entity.Reservation) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *Reservation) FindByID(ctx context.Context, id entity.ReservationID) (*entity.Reservation, error) {
	res := &entity.Reservation{}
	if err := t.db.NewSelect().Model(res).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Reservation) FindByUserIDAndTravelSpotID(ctx context.Context, userID entity.UserID, travelSpotID entity.TravelSpotID) (entity.Reservations, error) {
	res := entity.Reservations{}
	if err := t.db.NewSelect().Model(&res).
		Where("user_id = ?", userID).
		Where("travel_spot_id = ?", travelSpotID).
		Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Reservation) FindSpecialOfferByUserID(ctx context.Context, userID entity.UserID) (*entity.Reservation, error) {
	res := &entity.Reservation{}
	if err := t.db.NewSelect().Model(res).
		Where("user_id = ?", userID).
		Where("is_offer = 1").
		Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Reservation) GetEndedReservations(ctx context.Context, userID entity.UserID) (entity.Reservations, error) {
	res := entity.Reservations{}
	if err := t.db.NewSelect().Model(&res).
		Where("user_id = ?", userID).
		Where("(to_date <= CURDATE() OR is_offer = 1)").
		Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Reservation) GetUpcomingReservations(ctx context.Context, userID entity.UserID) (entity.Reservations, error) {
	res := entity.Reservations{}
	if err := t.db.NewSelect().Model(&res).
		Where("user_id = ?", userID).
		Where("from_date >= CURDATE()").Where("is_offer = 0").Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
