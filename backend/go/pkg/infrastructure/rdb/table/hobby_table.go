package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type Hobby struct {
	db *bun.DB
}

func NewHobby(db *bun.DB) *Hobby { return &Hobby{db: db} }

func (t *Hobby) Create(ctx context.Context, ents ...*entity.Hobby) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *Hobby) GetAll(ctx context.Context) (entity.Hobbies, error) {
	res := entity.Hobbies{}
	if err := t.db.NewSelect().Model(&res).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Hobby) GetByIDs(ctx context.Context, ids []entity.HobbyID) (entity.Hobbies, error) {
	res := entity.Hobbies{}
	if err := t.db.NewSelect().Model(&res).Where("id IN (?)", bun.In(ids)).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Hobby) GetByID(ctx context.Context, id entity.HobbyID) (*entity.Hobby, error) {
	res := &entity.Hobby{}
	if err := t.db.NewSelect().Model(res).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
