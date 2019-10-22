package main

import "fmt"

func getWeek() []string {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	week = append(week[6:], week[:6]...)
	return week
}

func main() {
	fmt.Println(getWeek())
}
