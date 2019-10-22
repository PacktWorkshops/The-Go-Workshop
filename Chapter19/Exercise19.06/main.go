package main
import (
	"crypto/rand"
	"fmt"
	"math/big"
	math "math/rand"
)
func main() {
	fmt.Println("Crypto random")
	for i := 1; i <= 10; i++ {
		data, _ := rand.Int(rand.Reader, big.NewInt(1000))
		fmt.Println(data)
	}
	fmt.Println("Math random")
	for i := 1; i <= 10; i++ {
		fmt.Println(math.Intn(1000))
	}
}
