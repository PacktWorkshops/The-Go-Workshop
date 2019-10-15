package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvHdrCol(hdr)
	fmt.Println("Result:")
	fmt.Println(result)
	fmt.Println()

	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	result2 := csvHdrCol(hdr2)
	fmt.Println("Result2:")
	fmt.Println(result2)
	fmt.Println()
}
func csvHdrCol(hdr []string) map[int]string {
	csvIdxToCol := make(map[int]string)
	for i, v := range hdr {
		v = strings.TrimSpace(v)
		switch strings.ToLower(v) {
		case "employee":
			csvIdxToCol[i] = v
		case "hours worked":
			csvIdxToCol[i] = v
		case "hourly rate":
			csvIdxToCol[i] = v
		}
	}
	return csvIdxToCol
}
