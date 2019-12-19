package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
)

type UserClient struct {
	ID   string
	Name string
}

type TxClient struct {
	ID          string
	User        *UserClient
	AccountFrom string
	AccountTo   string
	Amount      float64
}

type UserServer struct {
	ID string
}

type TxServer struct {
	ID          string
	User        UserServer
	AccountFrom string
	AccountTo   string
	Amount      *float32
}

func main() {
	// Create a dummy network
	var net bytes.Buffer
	// Create some data to encode
	clientTx := &TxClient{
		ID: "123456789",
		User: &UserClient{
			ID:   "ABCDEF",
			Name: "James",
		},
		AccountFrom: "Bob",
		AccountTo:   "Jane",
		Amount:      9.99,
	}
	// Encode the data, this writes it to our dummy network
	enc := gob.NewEncoder(&net)
	if err := enc.Encode(clientTx); err != nil {
		log.Fatal("error encoding: ", err)
	}
	serverTx, err := sendToServer(&net)
	if err != nil {
		log.Fatal("server error: ", err)
	}
	// What did we get?
	fmt.Printf("%#v\n", serverTx)
}

func sendToServer(net io.Reader) (*TxServer, error) {
	// Can't use a nil pointer
	tx := &TxServer{}
	// Decode
	dec := gob.NewDecoder(net)
	err := dec.Decode(tx)
	return tx, err
}
