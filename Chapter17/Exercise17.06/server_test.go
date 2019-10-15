package exercise6

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestAwesomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://awesome.com", nil)
	w := httptest.NewRecorder()
	AwesomeHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != "Welcome" {
		t.Errorf("We expected the string 'Welcome' but we received %s", string(body))
	}
}