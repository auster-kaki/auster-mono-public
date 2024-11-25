package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type Vendor struct {
	db *bun.DB
}

func NewVendor(db *bun.DB) *Vendor {
	return &Vendor{db: db}
}

func (t *Vendor) Create(ctx context.Context, ents ...entity.Vendor) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *Vendor) GetAll(ctx context.Context) (entity.Vendors, error) {
	res := entity.Vendors{}
	if err := t.db.NewSelect().Model(&res).Scan(ctx, &res); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Vendor) FindByID(ctx context.Context, id entity.VendorID) (*entity.Vendor, error) {
	res := entity.Vendor{}
	if err := t.db.NewSelect().Model(&res).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return &res, nil
}
