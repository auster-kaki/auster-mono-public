package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type EncounterID string

type Encounter struct {
	ID          EncounterID
	Name        string
	Place       string
	UserID      UserID
	Date        time.Time
	Description string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:encounter,alias:e"`
}

type Encounters []*Encounter
