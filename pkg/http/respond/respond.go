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

func HTTP(resp Response) {
	resp.Writer.Header().Set("Content-Type", "application/json")

	if resp.Err != nil {
		HTTPError(resp.Writer, resp.StatusCode, resp.Err)
		return
	}

	resp.Writer.WriteHeader(resp.StatusCode)
	json.NewEncoder(resp.Writer).Encode(resp.Body)
}

func HTTPError(w http.ResponseWriter, statusCode int, err error) {
	log.Println(err)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorBody{Message: err.Error()})
}
