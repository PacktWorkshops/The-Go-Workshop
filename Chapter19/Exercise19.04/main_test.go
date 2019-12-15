package main

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	const password = "mypassword"
	const key = "mysecurepassword"
	encrypted, err := encrypt([]byte(password), key)
	if err != nil {
		t.Error(err)
	}
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		t.Error(err)
	}
	if bytes.Equal(encrypted, decrypted) {
		t.Error("decrypted text does not mactch original value")
	}
}
