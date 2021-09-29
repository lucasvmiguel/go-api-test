package ping

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h := Handler{}
	h.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	statusCodeExpected := http.StatusOK
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	responseBodyExpected := "pong"
	if string(data) != responseBodyExpected {
		t.Errorf("expected %s got %v", responseBodyExpected, string(data))
	}
}
