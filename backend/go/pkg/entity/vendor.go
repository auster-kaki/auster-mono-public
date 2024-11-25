package entity

import (
	"github.com/uptrace/bun"
)

type VendorID string

type Vendor struct {
	ID      VendorID
	Name    string
	Address string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:vendor,alias:v"`
}

type Vendors []*Vendor
