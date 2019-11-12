package main
import (
	"fmt"
	"reflect"
)
type Animal struct {
	Name string
}
type Object struct {
	Type string
}
type Person struct {
	Name string
}
func MyPrint(input interface{}) {
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)
	switch {
	case t.Name() == "Animal":
		fmt.Println("I am a ", v.FieldByName("Name"))
	case t.Name() == "Object":
		fmt.Println("I am a ", v.FieldByName("Type"))
	default:
		fmt.Println("I got an unknown entity")
	}
}
func main() {
	table := Object{Type: "Chair"}
	MyPrint(table)
	tiger := Animal{Name: "Tiger"}
	MyPrint(tiger)
	gobin := Person{Name: "Gobin"}
	MyPrint(gobin)
}
