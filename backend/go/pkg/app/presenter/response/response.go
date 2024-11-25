package response

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/logging"
)

// OK - 200
func OK(w http.ResponseWriter, body any) {
	writeResponseJSON(w, http.StatusOK, body)
}

// Created - 201
func Created(w http.ResponseWriter, body any) {
	writeResponseJSON(w, http.StatusCreated, body)
}

// NoContent - 204
func NoContent(w http.ResponseWriter) {
	writeResponseJSON(w, http.StatusNoContent, nil)
}

// HandleError - 500
func HandleError(ctx context.Context, w http.ResponseWriter, err error) {
	logging.Error(ctx, fmt.Sprintf("code: %d, message: %s", http.StatusInternalServerError, err.Error()))
	var (
		code = http.StatusInternalServerError
		msg  = err.Error()
	)
	switch {
	case errors.Is(err, repository.ErrNotFound):
		code = http.StatusNotFound
		msg = "見つかりませんでした"
	case errors.Is(err, repository.ErrAlreadyExists):
		code = http.StatusConflict
		msg = "既に存在します"
	case errors.Is(err, repository.ErrDuplicate):
		code = http.StatusConflict
		msg = "重複しています"
	case errors.Is(err, repository.ErrNotExists):
		code = http.StatusNotFound
		msg = "存在しません"
	case errors.Is(err, repository.ErrUnimplemented):
		code = http.StatusNotImplemented
		msg = "未実装です"
	}
	writeResponseJSON(w, code, map[string]any{"message": msg})
}

func writeResponseJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, fmt.Errorf("failed to encode json. %w", err).Error(), http.StatusInternalServerError)
	}
}
