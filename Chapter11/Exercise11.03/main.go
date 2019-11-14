package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonData := []byte(`
{
    "id": 2,
    "lname": "Washington",
    "fname": "Bill",
    "IsEnrolled": true,
    "grades":[100,76,93,50],
    "class": 
        {
            "coursename": "World Lit",
            "coursenum": 101,
            "coursehours": 3
        }
}	
`)
	if !json.Valid(jsonData) {
		fmt.Printf("JSON is not valid: %s", jsonData)
		os.Exit(1)
	}
	var v interface{}
	err := json.Unmarshal(jsonData, &v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := v.(map[string]interface{})
	for k, v := range data {
		switch value := v.(type) {
		case string:
			fmt.Println("(string):", k, value)
		case float64:
			fmt.Println("(float64):", k, value)
		case bool:
			fmt.Println("(bool):", k, value)
		case []interface{}:
			fmt.Println("(slice):", k)
			for i, j := range value {
				fmt.Println("    ", i, j)
			}
		default:
			fmt.Println("(unknown):", k, value)
		}
	}
}
