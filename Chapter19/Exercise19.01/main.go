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
func area(input interface{}) float64 {
	inputType := reflect.TypeOf(input)
	if inputType.Name() == "circle" {
		val := reflect.ValueOf(input)
		radius := val.FieldByName("radius")
		return math.Pi * math.Pow(radius.Float(), 2)
	}
	if inputType.Name() == "rectangle" {
		val := reflect.ValueOf(input)
		length := val.FieldByName("length")
		breadth := val.FieldByName("breadth")
		return length.Float() * breadth.Float()
	}
	return 0
}
func main() {
	fmt.Printf("area of circle with radius 3 is : %f\n", area(circle{radius: 3}))
	fmt.Printf("area of rectangle with length 3 and breadth 7 is : %f\n", area(rectangle{length: 3, breadth: 7}))
}
