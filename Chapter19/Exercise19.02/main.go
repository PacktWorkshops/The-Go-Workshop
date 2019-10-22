package main
import (
	"fmt"
	"html/template"
	"net/http"
)
var content = `<html>
<head>
<title>My Blog</title>
</head>
<body>
   <h1>My Blog Site</h1>
   <h2> Comments </h2>
   {{.Comment}}
  <form action="/" method="post">
      Add Comment:<input type="text" name="input">
      <input type="submit" value="Submit">
  </form>
</body>
</html>`
type input struct {
	Comment string
}
func handler(w http.ResponseWriter, r *http.Request) {
	var userInput = &input{
		Comment: r.FormValue("input"),
	}
	t := template.Must(template.New("test").Parse(content))
	err := t.Execute(w, userInput)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
