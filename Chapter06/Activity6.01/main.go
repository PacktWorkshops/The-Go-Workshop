package main

import (
	"errors"
	"fmt"
)

var ErrInvalidLastName = errors.New("invalid last name")
var ErrInvalidRoutingNum = errors.New("invalid routing number")

func main() {
	fmt.Println(ErrInvalidLastName)
	fmt.Println(ErrInvalidRoutingNum)

}
