package exercise5

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAwesomeClient_GetUsers(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}))

	defer ts.Close()
	cl := AwesomeClient{ts.URL}
	rsp, err := cl.GetUsers()
	if err != nil {
		t.Errorf("We were expecting no errors but we received %v", err)
	}

	msg, err := ioutil.ReadAll(rsp.Body)
	rsp .Body.Close()
	if err != nil {
		t.Errorf("We were expecting no errors but we received %v", err)
	}

	if string(msg) != "Hello" {
		t.Errorf("We expected the message 'Hello' but we received %s", string(msg))
	}
}