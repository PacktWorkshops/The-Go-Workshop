package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr)
	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	csvHdrCol(hdr2)

}
func csvHdrCol(header []string) {
	csvHeadersToColumnIndex := make(map[int]string)
	for i, v := range header {
		v = strings.TrimSpace(v)
		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}
	fmt.Println(csvHeadersToColumnIndex)
}
