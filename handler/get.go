package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) Get() http.HandlerFunc {

	type response struct {
		Id       int    `json:"id"`
		UserName string `json:"user_name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			h.respond(w, errors.New("valid id must provide in path"), http.StatusBadRequest)
			return
		}

		user, err := h.service.Get(r.Context(), id)
		if err != nil {
			h.respond(w, err, 0)
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

func (h Handler) GetAll() http.HandlerFunc {

	type response struct {
		Id       int    `json:"id"`
		UserName string `json:"user_name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		users, err := h.service.GetAll(r.Context())
		if err != nil {
			h.respond(w, err, 0)
			return
		}
		resp := []response{}
		for _, user := range users {
			rp := response{
				Id:       user.Id,
				UserName: user.UserName,
				Password: "**********",
				Email:    user.Email,
			}
			resp = append(resp, rp)
		}
		h.respond(w, resp, http.StatusOK)
	}
}
