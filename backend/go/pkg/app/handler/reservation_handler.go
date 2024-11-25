package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type ReservationHandler struct {
	reservationUseCase *usecase.ReservationUseCase
}

func NewReservationHandler(r *usecase.ReservationUseCase) []Input {
	h := &ReservationHandler{reservationUseCase: r}
	return []Input{
		{method: http.MethodPost, path: "/reservations", handler: h.Create},
		{method: http.MethodGet, path: "/reservations", handler: h.List},
		{method: http.MethodGet, path: "/reservations/{reservation_id}", handler: h.GetReservation},
		{method: http.MethodPatch, path: "/reservations/{reservation_id}/diary_photo/users/{user_id}", handler: h.UpdateDiaryPhoto},
		{method: http.MethodPatch, path: "/reservations/{reservation_id}/diary_description", handler: h.UpdateDescription},
	}
}

func (h *ReservationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req request.Reservation
		ctx = r.Context()
	)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	from, err := time.ParseInLocation(time.DateOnly, req.FromDate, time.Local)
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	to, err := time.ParseInLocation(time.DateOnly, req.ToDate, time.Local)
	if err != nil {
		response.HandleError(ctx, w, err)
	}

	out, err := h.reservationUseCase.Create(ctx, &usecase.CreateReservationInput{
		UserID:            entity.UserID(req.UserID),
		TravelSpotID:      entity.TravelSpotID(req.TravelSpotID),
		TravelSpotDiaryID: entity.TravelSpotDiaryID(req.TravelSpotDiaryID),
		From:              from,
		To:                to,
	})
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.Created(w, out)
}

func (h *ReservationHandler) List(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		userID = r.URL.Query().Get("user_id")
		filter = r.URL.Query().Get("filter")
	)

	out, err := h.reservationUseCase.List(ctx, usecase.ListOutputInput{
		UserID: entity.UserID(userID),
		Status: filter,
	})
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	res := map[string]any{
		"reservations": []any{},
	}
	for _, reservation := range out.Reservations {
		title := ""
		description := ""
		if ts, ok := out.TravelSpotByID[reservation.TravelSpotID]; ok {
			title = ts.Title
			description = ts.Description
		}

		photoPath := ""
		if diary, ok := out.DiaryByID[reservation.TravelSpotDiaryID]; ok {
			photoPath = diary.PhotoPath
		}

		res["reservations"] = append(res["reservations"].([]any), map[string]any{
			"id":                      reservation.ID,
			"from_date":               reservation.FromDate,
			"to_date":                 reservation.ToDate,
			"is_offer":                reservation.IsOffer,
			"travel_spot_title":       title,
			"travel_spot_description": description,
			"diary_photo_path":        photoPath,
		})
	}
	response.OK(w, res)
}

func (h *ReservationHandler) GetReservation(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = entity.ReservationID(r.PathValue("reservation_id"))
	)

	out, err := h.reservationUseCase.GetReservation(ctx, id)
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	res := map[string]any{
		"id":                      out.Reservation.ID,
		"from_date":               out.Reservation.FromDate,
		"to_date":                 out.Reservation.ToDate,
		"is_offer":                out.Reservation.IsOffer,
		"travel_spot_title":       out.TravelSpot.Title,
		"travel_spot_description": out.TravelSpot.Description,
		"diary_title":             out.TravelSpotDiary.Title,
		"diary_description":       out.TravelSpotDiary.Description,
		"diary_photo_path":        out.TravelSpotDiary.PhotoPath,
		"travelSpotItineraries": func() []any {
			itineraries := make([]any, 0, len(out.TravelSpotItineraries))
			for _, itinerary := range out.TravelSpotItineraries {
				itineraries = append(itineraries, map[string]any{
					"id":          itinerary.ID,
					"title":       itinerary.Title,
					"description": itinerary.Description,
					"kind":        itinerary.Kind,
					"take_time":   itinerary.TakeTime,
					"price":       itinerary.Price,
					"order":       itinerary.Order,
					"items": func() []any {
						items := make([]any, 0, len(out.TravelSpotItineraryItemsByTravelSpotItineraryID[itinerary.ID]))
						for _, item := range out.TravelSpotItineraryItemsByTravelSpotItineraryID[itinerary.ID] {
							items = append(items, map[string]any{
								"id":     item.ID,
								"name":   item.Name,
								"number": item.Number,
							})
						}
						return items
					}(),
				})
			}
			return itineraries
		}(),
	}
	response.OK(w, res)
}

func (h *ReservationHandler) UpdateDiaryPhoto(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		id     = entity.ReservationID(r.PathValue("reservation_id"))
		userID = entity.UserID(r.PathValue("user_id"))
	)

	file, handler, err := r.FormFile("photo")
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	defer file.Close()

	photo, err := io.ReadAll(file)
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	out, err := h.reservationUseCase.UpdateDiaryPhoto(ctx, usecase.UpdateDiaryPhotoInput{
		ID:     id,
		UserID: userID,
		Photo: usecase.Photo{
			Filename:    handler.Filename,
			Body:        photo,
			ContentType: handler.Header.Get("Content-Type"),
		},
	})
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.Created(w, out)
}

func (h *ReservationHandler) UpdateDescription(w http.ResponseWriter, r *http.Request) {
	var (
		req request.DiaryDescription
		id  = entity.ReservationID(r.PathValue("reservation_id"))
		ctx = r.Context()
	)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	if err := h.reservationUseCase.UpdateDiaryDescription(ctx, usecase.UpdateDiaryDescriptionInput{
		ID:          id,
		UserID:      entity.UserID(req.UserID),
		Description: req.Description,
	}); err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.NoContent(w)
}
