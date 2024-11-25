package handler

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(u *usecase.UserUseCase) []Input {
	h := &UserHandler{userUseCase: u}
	return []Input{
		{method: http.MethodGet, path: "/users", handler: h.GetUsers},
		{method: http.MethodGet, path: "/users/{id}", handler: h.GetUser},
		{method: http.MethodPost, path: "/users", handler: h.CreateUser},
		{method: http.MethodPatch, path: "/users/{id}", handler: h.UpdateUser},
		{method: http.MethodGet, path: "/hobbies", handler: h.GetHobbies},
	}
}

func (h *UserHandler) GetHobbies(w http.ResponseWriter, r *http.Request) {
	out, err := h.userUseCase.GetHobbies(r.Context())
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	response.OK(w, out)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	out, err := h.userUseCase.GetUsers(r.Context())
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	response.OK(w, out)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	out, err := h.userUseCase.GetUser(r.Context(), entity.UserID(id))
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	// 画像ファイルとjsonを返す処理
	w.Header().Set("Content-Type", "multipart/mixed; boundary=boundary")

	// マルチパートレスポンスを生成
	multipartWriter := multipart.NewWriter(w)
	if err := multipartWriter.SetBoundary("boundary123"); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	// ユーザー情報をjsonで返す
	userJSON, err := json.Marshal(out.User)
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	// ユーザー情報をマルチパートレスポンスに書き込む
	userPart, err := multipartWriter.CreatePart(map[string][]string{
		"Content-Type": {"application/json"},
	})
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	if _, err := userPart.Write(userJSON); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	// 画像をマルチパートレスポンスに書き込む
	imagePart, err := multipartWriter.CreatePart(map[string][]string{
		"Content-Type": {"image/jpeg"},
	})
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	if _, err := imagePart.Write(out.Photo); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	// マルチパートレスポンスを閉じる
	if err := multipartWriter.Close(); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req request.User
	req.Name = r.FormValue("name")
	req.Gender = r.FormValue("gender")

	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	req.Age = age

	if err := json.Unmarshal([]byte(r.FormValue("hobbies")), &req.Hobbies); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	file, handler, err := r.FormFile("photo")
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	defer file.Close()

	photo, err := io.ReadAll(file)
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	id, err := h.userUseCase.CreateUser(r.Context(), &usecase.UserInput{
		Name:   req.Name,
		Gender: req.Gender,
		Age:    req.Age,
		Hobbies: func() entity.Hobbies {
			hobbies := make(entity.Hobbies, len(req.Hobbies))
			for i, hobby := range req.Hobbies {
				hobbies[i] = &entity.Hobby{
					ID:   entity.HobbyID(hobby.ID),
					Name: hobby.Name,
				}
			}
			return hobbies
		}(),
		Photo: usecase.Photo{
			Filename:    handler.Filename,
			Body:        photo,
			ContentType: handler.Header.Get("Content-Type"),
		},
	})
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	response.Created(w, map[string]string{"id": string(id)})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req request.User
	req.Name = r.FormValue("name")
	req.Gender = r.FormValue("gender")

	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	req.Age = age

	if err := json.Unmarshal([]byte(r.FormValue("hobbies")), &req.Hobbies); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	file, handler, err := r.FormFile("photo")
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	defer file.Close()

	photo, err := io.ReadAll(file)
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	if err := h.userUseCase.UpdateUser(r.Context(), &usecase.UserInput{
		ID:     entity.UserID(r.PathValue("id")),
		Name:   req.Name,
		Gender: req.Gender,
		Age:    req.Age,
		Hobbies: func() entity.Hobbies {
			hobbies := make(entity.Hobbies, len(req.Hobbies))
			for i, hobby := range req.Hobbies {
				hobbies[i] = &entity.Hobby{
					ID:   entity.HobbyID(hobby.ID),
					Name: hobby.Name,
				}
			}
			return hobbies
		}(),
		Photo: usecase.Photo{
			Filename:    handler.Filename,
			Body:        photo,
			ContentType: handler.Header.Get("Content-Type"),
		},
	}); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	response.OK(w, nil)
}
