package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPageWithCounter_ServeHTTP(t *testing.T) {
	cntr := PageWithCounter{heading: "title", content:"some content"}

	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	w := httptest.NewRecorder()

	cntr.ServeHTTP(w,r)
	if cntr.counter != 1 {
		t.Errorf("We expected 1 view but we received %d", cntr.counter)
	}

	msg := fmt.Sprintf("<h1>%s</h1>\n<p>%s<p>\n<p>Views: %d</p>", cntr.heading, cntr.content, 1)

	if w.Body.String() != msg {
		t.Errorf("Expected  '%s' but received %s", msg, w.Body.String())
	}

	w = httptest.NewRecorder()

	cntr.ServeHTTP(w,r)
	if cntr.counter != 2 {
		t.Errorf("We expected 1 view but we received %d", cntr.counter)
	}

	msg = fmt.Sprintf("<h1>%s</h1>\n<p>%s<p>\n<p>Views: %d</p>", cntr.heading, cntr.content, 2)

	if w.Body.String() != msg {
		t.Errorf("Expected  '%s' but received %s", msg, w.Body.String())
	}
}
