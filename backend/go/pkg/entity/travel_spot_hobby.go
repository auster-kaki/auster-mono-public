package entity

import "github.com/uptrace/bun"

type TravelSpotHobby struct {
	TravelSpotID TravelSpotID
	HobbyID      HobbyID

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:travel_spot_hobby,alias:tsh"`
}

type TravelSpotHobbies []*TravelSpotHobby

func (t TravelSpotHobbies) TravelSpotIDs() []TravelSpotID {
	ids := make([]TravelSpotID, len(t))
	for i, v := range t {
		ids[i] = v.TravelSpotID
	}
	return ids
}
