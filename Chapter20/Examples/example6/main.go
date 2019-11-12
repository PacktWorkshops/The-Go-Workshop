package main
import (
	"fmt"
	"reflect"
)
func main() {
	runDeepEqual(nil, nil)
	runDeepEqual(make([]int, 10), make([]int, 10))
	runDeepEqual([3]int{1, 2, 3}, [3]int{1, 2, 3})
	runDeepEqual(map[int]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"})
}
func runDeepEqual(a, b interface{}) {
	fmt.Printf("%v DeepEqual %v : %v\n", a, b, reflect.DeepEqual(a, b))
}
