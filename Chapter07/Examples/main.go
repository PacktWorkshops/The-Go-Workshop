package main
import (
  "encoding/json"
  "fmt"
  "io"
  "strings"
)
type Person struct {
  Name string `json:"name"`
  Age  int    `json:"age"`
}
func main() {
  s := `{"Name":"Joe","Age":18}`
  s2 := `{"Name":"Jane","Age":21}`
  p, err := loadPerson(strings.NewReader(s))
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(p)
  p2, err := loadPerson2(s2)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(p2)
}
func loadPerson2(s string) (Person, error) {
  var p Person
  err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
  if err != nil {
    return p, err
  }
  return p, nil
}
func loadPerson(r io.Reader) (Person, error) {
  var p Person
  err := json.NewDecoder(r).Decode(&p)
  if err != nil {
    return p, err
  }
  return p, err
}