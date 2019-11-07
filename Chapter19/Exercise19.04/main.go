package main
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)
func main() {
	const key = "mysecurepassword"
	encrypted, err := encrypt([]byte("Hello World!"), key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Encrypted Text: %x\n", encrypted)
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Decrypted Text: ", string(decrypted))
}
func encrypt(data []byte, key string) (resp []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}
	var dst []byte
	nonce := make([]byte, gcm.NonceSize())
	dst = nonce
	if _, err := rand.Read(nonce); err != nil {
		return resp, err
	}
	return gcm.Seal(dst, nonce, data, []byte("test")), nil
}
func decrypt(data []byte, key string) (resp []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}
	ciphertext := data[gcm.NonceSize():]
	nonce := data[:gcm.NonceSize()]
	resp, err = gcm.Open(nil, nonce, ciphertext, []byte("test"))
	if err != nil {
		return resp, fmt.Errorf("error decrypting data: %v", err)
	}
	return resp, nil
}
