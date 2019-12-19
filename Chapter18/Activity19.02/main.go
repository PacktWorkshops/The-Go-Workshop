package main
import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"os"
	"time"
)
func main() {
	// Generate CA certificates
	caCert, caPriv, err := generateCert("CA cert", nil, nil)
	if err != nil {
		fmt.Printf("error generating server certificate: %v", err)
		os.Exit(1)
	} else {
		fmt.Println("ca certificate generated successfully")
	}
	// User CA cert to generate and sign server certificate
	cert, _, err := generateCert("Test Cert", caCert, caPriv)
	if err != nil {
		fmt.Printf("error generating server certificate: %v", err)
		os.Exit(1)
	} else {
		fmt.Println("leaf certificate generated successfully")
	}
	rootCerts := x509.NewCertPool()
	rootCerts.AddCert(caCert)
	options := x509.VerifyOptions{
		Roots: rootCerts,
	}
	_, err = cert.Verify(options)
	if err != nil {
		fmt.Printf("error verifying certificate: %v", err)
		os.Exit(1)
	} else {
		fmt.Println("leaf certificate successfully verified")
	}
}
func generateCert(cn string, caCert *x509.Certificate, caPriv crypto.PrivateKey) (cert *x509.Certificate, privateKey crypto.PrivateKey, err error) {
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		return cert, privateKey, err
	}
	var isCA bool
	if caCert == nil {
		isCA = true
	}
	template := &x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               pkix.Name{CommonName: cn},
		NotBefore:             time.Now().Add(-1 * time.Hour),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA: isCA,
	}
	ecsdaKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return cert, nil, err
	}
	if caCert == nil {
		caCert = template
		caPriv = ecsdaKey
	}
	DER, err := x509.CreateCertificate(rand.Reader, template, caCert, ecsdaKey.Public(), caPriv)
	if err != nil {
		return cert, ecsdaKey, err
	}
	cert, err = x509.ParseCertificate(DER)
	if err != nil {
		fmt.Printf("error parsing x509 certificate: %v", err)
		return cert, ecsdaKey, err
	}
	return cert, ecsdaKey, nil
}
