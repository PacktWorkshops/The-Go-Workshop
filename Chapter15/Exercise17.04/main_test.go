package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloFullDetails(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, `/?name=john&surname=smith&age=30&id=1`, strings.NewReader(""))
	w := httptest.NewRecorder()

	var testStr = `
<html>
  <h1>Customer 1</h1>

    <p>Details:</p>
    <ul>
        <li>Name: john</li>
        <li>Surname: smith</li>
        <li>Age: 30</li>
    </ul>

</html>
`

	Hello(w, r)

	rpl := strings.NewReplacer(" ", "", "\n", "")

	if rpl.Replace(w.Body.String()) != rpl.Replace(testStr, ) {
		t.Errorf("Wrong text, expected %s but received %s", rpl.Replace(testStr), rpl.Replace(w.Body.String()))
	}
}

func TestHelloMissingSurname(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, `/?name=john&age=30&id=1`, strings.NewReader(""))
	w := httptest.NewRecorder()

	var testStr = `
<html>
  <h1>Customer 1</h1>

  <p>Details:</p>
  <ul>
      <li>Name: john</li>

      <li>Age: 30</li>
  </ul>

</html>
`

	Hello(w, r)

	rpl := strings.NewReplacer(" ", "", "\n", "")

	if rpl.Replace(w.Body.String()) != rpl.Replace(testStr, ) {
		t.Errorf("Wrong text, expected %s but received %s", rpl.Replace(testStr), rpl.Replace(w.Body.String()))
	}
}

func TestHelloMissingID(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, `/?name=john&age=30`, strings.NewReader(""))
	w := httptest.NewRecorder()

	var testStr = `
<html>
   <h1>Customer 0</h1>
  <p>Data not available</p>
</html>
`

	Hello(w, r)

	rpl := strings.NewReplacer(" ", "", "\n", "")

	if rpl.Replace(w.Body.String()) != rpl.Replace(testStr, ) {
		t.Errorf("Wrong text, expected %s but received %s", rpl.Replace(testStr), rpl.Replace(w.Body.String()))
	}
}
