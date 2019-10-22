package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPageWithCounter_ServeHTTP(t *testing.T) {
	cntr := PageWithCounter{Heading: "title", Content:"some Content"}

	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	w := httptest.NewRecorder()

	cntr.ServeHTTP(w,r)
	if cntr.Counter != 1 {
		t.Errorf("We expected 1 view but we received %d", cntr.Counter)
	}

	msg := fmt.Sprintf("{\"views\":%d,\"title\":\"%s\",\"content\":\"%s\"}", 1, cntr.Heading, cntr.Content)

	if w.Body.String() != msg {
		t.Errorf("Expected  '%s' but received %s", msg, w.Body.String())
	}

	w = httptest.NewRecorder()

	cntr.ServeHTTP(w,r)
	if cntr.Counter != 2 {
		t.Errorf("We expected 1 view but we received %d", cntr.Counter)
	}

	msg = fmt.Sprintf("{\"views\":%d,\"title\":\"%s\",\"content\":\"%s\"}", 2, cntr.Heading, cntr.Content)

	if w.Body.String() != msg {
		t.Errorf("Expected  '%s' but received %s", msg, w.Body.String())
	}
}