package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type User struct {
	db *bun.DB
}

func NewUser(db *bun.DB) *User {
	return &User{db: db}
}

func (t *User) GetAll(ctx context.Context) (entity.Users, error) {
	res := entity.Users{}
	if err := t.db.NewSelect().Model(&entity.User{}).Scan(ctx, &res); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *User) FindByID(ctx context.Context, id entity.UserID) (*entity.User, error) {
	user := &entity.User{}
	if err := t.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return user, nil
}

func (t *User) Create(ctx context.Context, ents ...*entity.User) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *User) Update(ctx context.Context, user *entity.User) error {
	if _, err := t.db.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}
