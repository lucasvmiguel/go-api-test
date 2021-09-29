package client

import "net/http"

type HTTP interface {
	Do(*http.Request) (*http.Response, error)
}
