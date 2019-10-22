package main

import "fmt"

func getData() []interface{} {
	return []interface{}{
		1,
		3.14,
		"hello",
		true,
		struct{}{},
	}
}

func getTypeName(v interface{}) string {
	switch v.(type) {
	case int, int32, int64:
		return "int"
	case float64, float32:
		return "float"
	case bool:
		return "bool"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func main() {
	data := getData()
	for i := 0; i < len(data); i++ {
		fmt.Printf("%v is %v\n", data[i], getTypeName(data[i]))
	}
}
