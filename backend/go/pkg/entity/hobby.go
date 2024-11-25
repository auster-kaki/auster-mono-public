package entity

import "github.com/uptrace/bun"

type HobbyID string

type Hobby struct {
	ID   HobbyID
	Name string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:hobby,alias:h"`
}

type Hobbies []*Hobby

func (e *Hobbies) Names() []string {
	names := make([]string, len(*e))
	for i, hobby := range *e {
		names[i] = hobby.Name
	}
	return names
}
