package main

import (
	"crypto/tls"
	"testing"
)

func TestGenerate(t *testing.T) {
	cert, key, err := generate()
	if err != nil {
		t.Errorf("error generating client certificate: %v", err)
	}
	_, err = tls.X509KeyPair(cert, key)
	if err != nil {
		t.Errorf("invalid key pair: %v", err)
	}
}
