package main
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)
func main() {
	serverCert, serverKey, err := generate()
	if err != nil {
		fmt.Printf("error generating server certificate: %v", err)
		os.Exit(1)
	}
	ioutil.WriteFile("private.key", serverKey, 0600)
	ioutil.WriteFile("cert.pem", serverCert, 0777)
	clientCert, clientKey, err := generate()
	if err != nil {
		fmt.Printf("error generating client certificate: %v", err)
		os.Exit(1)
	}
	ioutil.WriteFile("client.key", clientKey, 0600)
	ioutil.WriteFile("client_cert.pem", clientCert, 0777)
	clientTLSCert, err := tls.X509KeyPair(clientCert, clientKey)
	if err != nil {
		log.Fatalf("invalid key pair: %v", err)
	}
	errChan := make(chan error, 2)
	go func() {
		for {
			errChan <- runServer("cert.pem", "private.key", clientCert)
		}
	}()
	time.Sleep(time.Second * 3)
	go func() {
		for {
			time.Sleep(time.Second * 3)
			err = client(serverCert, clientTLSCert)
			if err != nil {
				errChan <- err
			}
		}
	}()
	fmt.Println(<-errChan)
}
func generate() (cert []byte, privateKey []byte, err error) {
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		return cert, privateKey, err
	}
	notBefore := time.Now()
	ca := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"example.com"},
		},
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(time.Hour * 24 * 365),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return cert, privateKey, err
	}
	DER, err := x509.CreateCertificate(rand.Reader, ca, ca, &rsaKey.PublicKey, rsaKey)
	if err != nil {
		return cert, privateKey, err
	}
	b := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: DER,
	}
	cert = pem.EncodeToMemory(&b)

	privateKey = pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(rsaKey),
		})
	return cert, privateKey, nil
}
func client(caCert []byte, ClientCert tls.Certificate) (err error) {
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      certPool,
				Certificates: []tls.Certificate{ClientCert},
			},
		},
	}
	resp, err := client.Get("https://127.0.0.1:443")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s: %s", time.Now().Format(time.Stamp), body)
	return err
}
func runServer(certFile string, key string, clientCert []byte) (err error) {
	fmt.Println("starting HTTP server")
	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr:    ":443",
		Handler: nil,
	}
	cert, err := tls.LoadX509KeyPair(certFile, key)
	if err != nil {
		return err
	}
	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCert)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCertPool,
	}
	if err != nil {
		return err
	}
	conn, err := net.Listen("tcp", server.Addr)
	listener := tls.NewListener(conn, tlsConfig)
	err = server.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: Ping\n", time.Now().Format(time.Stamp))
	fmt.Fprintf(w, "Pong\n")
}
