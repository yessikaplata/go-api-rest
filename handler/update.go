package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yessikaplata/go-api-rest/service"
)

func (h Handler) Update() http.HandlerFunc {

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
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			h.respond(w, errors.New("valid id must provide in path"), http.StatusBadRequest)
			return
		}

		req := request{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			h.respond(w, err, http.StatusBadRequest)
			return
		}

		user, err := h.service.Update(r.Context(), service.UpdateParams{
			Id:       id,
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
