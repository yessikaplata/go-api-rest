package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yessikaplata/go-api-rest/service"
)

func (h Handler) Create() http.HandlerFunc {

	// contract
	type request struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type response struct {
		Id       int    `json:"id"`
		UserName string `json:"user_name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			h.respond(w, err, http.StatusBadRequest)
			return
		}

		user, err := h.service.Create(r.Context(), service.CreateParams{
			UserName: req.UserName,
			Password: req.Password,
			Email:    req.Email,
		})

		if err != nil {
			h.respond(w, err, http.StatusBadRequest)
			return
		}

		resp := response{
			Id:       user.Id,
			UserName: user.UserName,
			Password: "**********",
			Email:    user.Email,
		}
		h.respond(w, resp, http.StatusOK)
	}
}
