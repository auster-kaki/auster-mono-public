package entity

import (
	"github.com/uptrace/bun"
)

type UserID string

type User struct {
	ID          UserID
	Name        string
	Gender      string
	Age         int
	ProfilePath string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:user,alias:u"`
}

type Users []*User
