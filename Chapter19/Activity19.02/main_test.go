package main

import (
	"bytes"
	"crypto/x509"
	"testing"
)

func TestGenerateCert(t *testing.T) {
	caCert, _, err := generateCert("CAAuthority", nil, nil)
	if err != nil {
		t.Error(err)
	}
	pool := x509.NewCertPool()
	pool.AddCert(caCert)
	subjects := pool.Subjects()
	var found bool
	for _, sub := range subjects {
		if bytes.Contains(sub, []byte("CAAuthority")) {
			found = true
			break
		}
	}
	if !found {
		t.Error("failed to verify ca subject name")
	}
}
