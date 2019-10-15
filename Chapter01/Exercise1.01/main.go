package main
import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(strings.Repeat("*", rand.Intn(5)+1))
}