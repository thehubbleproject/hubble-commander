package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ReadRESTReq reads and unmarshals a Request's body to the the BaseReq stuct.
// Writes an error response to ResponseWriter and returns true if errors occurred.
func ReadRESTReq(r *http.Request, req interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, req)
	return err
}

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
