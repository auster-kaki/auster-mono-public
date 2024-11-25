package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type Encounter struct {
	db *bun.DB
}

func NewEncounter(db *bun.DB) *Encounter {
	return &Encounter{db: db}
}

func (t *Encounter) Create(ctx context.Context, ents ...entity.Encounter) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *Encounter) Update(ctx context.Context, ent *entity.Encounter) error {
	if _, err := t.db.NewUpdate().Model(ent).Where("id = ?", ent.ID).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *Encounter) FindByID(ctx context.Context, id entity.EncounterID) (*entity.Encounter, error) {
	res := &entity.Encounter{}
	if err := t.db.NewSelect().Model(res).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Encounter) GetByUserID(ctx context.Context, userID entity.UserID) (entity.Encounters, error) {
	res := entity.Encounters{}
	if err := t.db.NewSelect().Model(&res).Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
