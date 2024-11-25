package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type TravelSpotHandler struct {
	travelSpotUseCase *usecase.TravelSpotUseCase
}

func NewTravelSpotHandler(u *usecase.TravelSpotUseCase) []Input {
	h := &TravelSpotHandler{travelSpotUseCase: u}
	return []Input{
		{method: http.MethodGet, path: "/travel_spots", handler: h.GetTravelSpots},
		{method: http.MethodPost, path: "/travel_spots/{travel_spot_id}/diary", handler: h.CreateDiary},
		{method: http.MethodGet, path: "/travel_spots/{travel_spot_id}/itineraries", handler: h.GetItineraries},
	}
}

func (h *TravelSpotHandler) GetTravelSpots(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.Background()
		userID  = r.URL.Query().Get("user_id")
		hobbyID = r.URL.Query().Get("hobby_id")
	)
	out, err := h.travelSpotUseCase.GetTravelSpots(ctx, entity.UserID(userID), entity.HobbyID(hobbyID))
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.OK(w, out)
}

func (h *TravelSpotHandler) CreateDiary(w http.ResponseWriter, r *http.Request) {
	var (
		ctx          = context.Background()
		travelSpotID = entity.TravelSpotID(r.PathValue("travel_spot_id"))
		req          request.Diary
	)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(ctx, w, fmt.Errorf("failed to decode request body: %w", err))
		return
	}

	date, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		response.HandleError(ctx, w, fmt.Errorf("failed to parse date: %w", err))
		return
	}
	out, err := h.travelSpotUseCase.CreateDiary(ctx, entity.UserID(req.UserID), travelSpotID, date)
	if err != nil {
		response.HandleError(ctx, w, fmt.Errorf("failed to create diary: %w", err))
		return
	}

	// Content-Typeをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// レスポンスデータの構造を修正
	res := map[string]any{
		"ID":        out.ID,
		"Title":     out.Title,
		"Body":      out.Description,
		"PhotoPath": out.PhotoPath,
	}

	// JSONエンコード
	if err := json.NewEncoder(w).Encode(res); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

}

func (h *TravelSpotHandler) GetItineraries(w http.ResponseWriter, r *http.Request) {
	var (
		ctx          = context.Background()
		userID       = entity.UserID(r.URL.Query().Get("user_id"))
		travelSpotID = entity.TravelSpotID(r.PathValue("travel_spot_id"))
	)
	out, err := h.travelSpotUseCase.GetItineraries(ctx, userID, travelSpotID)
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.OK(w, out)
}
