package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Decode[T any](r *http.Request, v *T) error {
	var (
		buf = new(bytes.Buffer)
		tee = io.TeeReader(r.Body, buf)
	)
	if err := json.NewDecoder(tee).Decode(v); err != nil {
		return fmt.Errorf("failed to decode request body:[%s] %w", buf.String(), err)
	}
	return nil
}

type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Gender  string  `json:"gender"`
	Hobbies []Hobby `json:"hobbies"`
}

type Hobby struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Encounter struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Place       string `json:"place"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type TravelSpot struct {
	UserID  string `json:"user_id"`
	HobbyID string `json:"hobby_id"`
}

type Diary struct {
	UserID          string `json:"user_id"`
	HobbyID         string `json:"hobby_id"`
	TravelSpotID    string `json:"travel_spot_id"`
	DestinationName string `json:"destination_name"`
	Date            string `json:"date"`
}

type Reservation struct {
	UserID            string `json:"user_id"`
	TravelSpotID      string `json:"travel_spot_id"`
	TravelSpotDiaryID string `json:"travel_spot_diary_id"`
	FromDate          string `json:"from_date"`
	ToDate            string `json:"to_date"`
}

type DiaryDescription struct {
	UserID      string `json:"user_id"`
	Description string `json:"description"`
}
