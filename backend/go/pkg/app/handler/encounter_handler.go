package handler

import (
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type EncounterHandler struct {
	encounterUseCase *usecase.EncounterUseCase
}

func NewEncounterHandler(e *usecase.EncounterUseCase) []Input {
	h := &EncounterHandler{encounterUseCase: e}
	return []Input{
		{method: http.MethodGet, path: "/encounters", handler: h.GetEncounters},
		{method: http.MethodPost, path: "/encounters", handler: h.Create},
		{method: http.MethodGet, path: "/encounters/{id}", handler: h.GetEncounter},
		{method: http.MethodPatch, path: "/encounters/{id}", handler: h.Update},
	}
}

func (h *EncounterHandler) GetEncounters(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	out, err := h.encounterUseCase.GetEncounters(r.Context(), entity.UserID(userID))
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	response.OK(w, out)
}

func (h *EncounterHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.Encounter
	if err := request.Decode(r, &req); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	if err := h.encounterUseCase.Create(r.Context(), &usecase.CrateEncounterInput{
		UserID:      entity.UserID(req.UserID),
		Name:        req.Name,
		Place:       req.Place,
		Date:        req.Date,
		Description: req.Description,
	}); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	response.Created(w, nil)
}

func (h *EncounterHandler) GetEncounter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	out, err := h.encounterUseCase.GetEncounter(r.Context(), entity.EncounterID(id))
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	response.OK(w, out)
}

func (h *EncounterHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req request.Encounter
	if err := request.Decode(r, &req); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	if err := h.encounterUseCase.Update(r.Context(), &usecase.UpdateEncounterInput{
		ID:          entity.EncounterID(r.PathValue("id")),
		Name:        req.Name,
		Place:       req.Place,
		Date:        req.Date,
		Description: req.Description,
	}); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	response.OK(w, nil)
}
