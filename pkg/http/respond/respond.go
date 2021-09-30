package respond

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Body       interface{}
	StatusCode int
	Err        error
	Writer     http.ResponseWriter
}

type errorBody struct {
	Message string `json:"message"`
}

var (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json"
)

func HTTP(resp Response) {
	if resp.Err != nil {
		HTTPError(resp.Writer, resp.StatusCode, resp.Err)
		return
	}

	resp.Writer.Header().Set(contentTypeKey, contentTypeValue)
	resp.Writer.WriteHeader(resp.StatusCode)
	json.NewEncoder(resp.Writer).Encode(resp.Body)
}

func HTTPError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set(contentTypeKey, contentTypeValue)

	log.Println(err)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorBody{Message: err.Error()})
}
