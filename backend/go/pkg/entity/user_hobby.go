package entity

import "github.com/uptrace/bun"

type UserHobby struct {
	UserID  UserID
	HobbyID HobbyID

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:user_hobby,alias:uh"`
}

type UserHobbies []*UserHobby

func (uh UserHobbies) HobbyIDs() []HobbyID {
	ids := make([]HobbyID, 0, len(uh))
	for _, h := range uh {
		ids = append(ids, h.HobbyID)
	}
	return ids
}

func (uh UserHobbies) UserIDs() []UserID {
	ids := make([]UserID, 0, len(uh))
	for _, h := range uh {
		ids = append(ids, h.UserID)
	}
	return ids
}

func (uh UserHobbies) UserIDMap() map[UserID]UserHobbies {
	m := make(map[UserID]UserHobbies)
	for _, h := range uh {
		m[h.UserID] = append(m[h.UserID], h)
	}
	return m
}
