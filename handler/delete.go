package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) Delete() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			h.respond(w, errors.New("valid id must provide in path"), http.StatusBadRequest)
			return
		}

		err = h.service.Delete(r.Context(), id)
		if err != nil {
			h.respond(w, err, 0)
			return
		}
		h.respond(w, nil, http.StatusOK)
	}
}
