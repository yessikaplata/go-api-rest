package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yessikaplata/go-api-rest/service/errors"
)

type ErrorResponse struct {
	ErrorMessage string
}

func (h Handler) respond(w http.ResponseWriter, data interface{}, status int) {
	var respData interface{}

	switch v := data.(type) {
	case nil:
	case errors.ArgumentsError:
		status = http.StatusBadRequest
		respData = ErrorResponse{ErrorMessage: v.Error()}
	case errors.UserNotFoundError:
		status = http.StatusNotFound
		respData = ErrorResponse{ErrorMessage: v.Error()}
	case error:
		if http.StatusText(status) == "" {
			status = http.StatusInternalServerError
		} else {
			respData = ErrorResponse{ErrorMessage: v.Error()}
		}
	default:
		respData = data
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		err := json.NewEncoder(w).Encode(respData)
		if err != nil {
			http.Error(w, "Could not encode in json", http.StatusBadRequest)
			return
		}
	}
}
