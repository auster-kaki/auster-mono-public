package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type UserHobby struct {
	db *bun.DB
}

func NewUserHobby(db *bun.DB) *UserHobby { return &UserHobby{db: db} }

func (t *UserHobby) Create(ctx context.Context, ents ...*entity.UserHobby) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *UserHobby) DeleteByUserID(ctx context.Context, userID entity.UserID) error {
	if _, err := t.db.NewDelete().Model(&entity.UserHobby{}).Where("user_id = ?", userID).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *UserHobby) GetByUserID(ctx context.Context, userID entity.UserID) (entity.UserHobbies, error) {
	res := entity.UserHobbies{}
	if err := t.db.NewSelect().Model(&res).Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
