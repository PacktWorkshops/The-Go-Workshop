package main
import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)
func main() {
	r := random(1, 20)
	err := a(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = b(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn((max-min)+1) + min
}
func a(i int) error {
	if i < 10 {
		fmt.Println("Error is in func a")
		return errors.New("Incorrect value")
	}
	return nil
}
func b(i int) error {
	if i >= 10 {
		fmt.Println("Error is in func b.")
		return errors.New("Incorrect value")
	}
	return nil
}
