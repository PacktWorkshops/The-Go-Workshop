package main
import (
  "fmt"
)
type cat struct {
  name string
}
func main() {
  c:=cat{name:"oreo"}
  i := []interface{}{42, "The book club", true,c}
  typeExample(i)
}
func typeExample(i []interface{}) {
  for _, x := range i {
    switch v := x.(type) {
    case int:
      fmt.Printf("%v is int\n", v)
    case string:
      fmt.Printf("%v is a string\n",v)
    case bool:
      fmt.Printf("a bool %v\n", v)
    default:
      fmt.Printf("Unknown type %T\n", v)
    }
  }
}