package table

import (
	"context"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/uptrace/bun"
)

type TravelSpotHobby struct {
	db *bun.DB
}

func NewTravelSpotHobby(db *bun.DB) *TravelSpotHobby { return &TravelSpotHobby{db: db} }

func (t *TravelSpotHobby) GetByHobbyID(ctx context.Context, hobbyID entity.HobbyID) (entity.TravelSpotHobbies, error) {
	res := entity.TravelSpotHobbies{}
	if err := t.db.NewSelect().Model(&res).Where("hobby_id = ?", hobbyID).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
