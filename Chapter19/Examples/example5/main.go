package main
import (
	"fmt"
	"reflect"
)
func main() {
	var x = 5
	Print(x)
	var y = []string{"test"}
	Print(y)
	var z = map[string]string{"a": "b"}
	Print(z)
}
func Print(a interface{}) {
	fmt.Println("Type: ", reflect.TypeOf(a))
	fmt.Println("Value: ", reflect.ValueOf(a))
}
