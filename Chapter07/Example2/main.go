package main
import (
  "fmt"
)
type Speaker interface {
  Speak() string
}
type cat struct {
  name string
}
func main() {
  c := cat{name: "oreo"}
  i := 99
  b := false
  str := "test"
  catDetails(c)
  emptyDetails(c)
  emptyDetails(i)
  emptyDetails(b)
  emptyDetails(str)
}
func (c cat) Speak() string {
  return "Purr Meow"
}
func emptyDetails(i interface{}) {
  fmt.Printf("(%v, %T)\n", i, i)
}
func catDetails(i Speaker) {
  fmt.Printf("(%v, %T)\n", i, i)
}