package exercise6

import (
	"io"
	"net/http"
)

func AwesomeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome")
}

