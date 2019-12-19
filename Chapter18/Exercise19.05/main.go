package main
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)
func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Printf("error generating rsa key: %v", err)
	}
	publicKey := privateKey.PublicKey
	text := []byte("My Secret Text")
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, text, nil)
	if err != nil {
		fmt.Printf("error encrypting data: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Encrypted cyphertext: %x\n", ciphertext)
	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		fmt.Printf("error decrypting data: %v", err)
		os.Exit(1)
	}
	fmt.Println("Decrypted text: ", string(decrypted))
}
