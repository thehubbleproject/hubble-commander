package rest

import (
	"encoding/json"
	"net/http"
)

// WriteRESTResp writes the response and the error to proper JSON format
func WriteRESTResp(w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	output, err := json.Marshal(resp)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

type ErrorResponse struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error"`
}

// NewErrorResponse creates a new ErrorResponse instance.
func NewErrorResponse(code int, err string) ErrorResponse {
	return ErrorResponse{Code: code, Error: err}
}

// WriteErrorResponse prepares and writes a HTTP error
// given a status code and an error message.
func WriteErrorResponse(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(MustMarshallJSON(NewErrorResponse(0, err)))
}

// MustMarshallJSON either marshalls to json or panics
func MustMarshallJSON(o interface{}) []byte {
	bz, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return bz
}
