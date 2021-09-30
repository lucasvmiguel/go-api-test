package ping

import (
	"net/http"
	"testing"

	"github.com/lucasvmiguel/go-api-test/pkg/test"
)

func TestPingHandler(t *testing.T) {
	h := &Handler{}
	res, resBody := test.HttpRequest(h, nil)
	defer res.Body.Close()

	statusCodeExpected := http.StatusOK
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	responseBodyExpected := "pong"
	if resBody != responseBodyExpected {
		t.Errorf("expected %s got %v", responseBodyExpected, resBody)
	}
}
