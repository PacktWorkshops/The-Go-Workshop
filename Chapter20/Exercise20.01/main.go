package main
import (
	"fmt"
	"math"
	"reflect"
)
type circle struct {
	radius float64
}
type rectangle struct {
	length  float64
	breadth float64
}
func area(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("type is:", inputType.Name())
	if inputType.Name() == "circle" {
		val := reflect.ValueOf(input)
		radius := val.FieldByName("radius")
		fmt.Println("area is: ", math.Pi*math.Pow(radius.Float(), 2))
	}
	if inputType.Name() == "rectangle" {
		val := reflect.ValueOf(input)
		length := val.FieldByName("length")
		breadth := val.FieldByName("breadth")
		fmt.Println("area is: ", length.Float()*breadth.Float())
	}
}
func main() {
	area(circle{radius: 3})
	area(rectangle{length: 3.1, breadth: 7.2})
}
