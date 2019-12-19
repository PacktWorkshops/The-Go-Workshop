package main
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
func main() {
	password := "mysecretpassword"
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	fmt.Println("Plain Text Password:", password)
	fmt.Println("Hashed Password:    ", string(encrypted))
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	if err == nil {
		fmt.Println("Password matched")
	}
}
