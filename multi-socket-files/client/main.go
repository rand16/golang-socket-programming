package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"

	"golang-socket-programming/pkg/message"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/sock1.sock"
)

var textsocket = map[string]string{
	"socket1": "/tmp/sock1.sock",
	"socket2": "/tmp/sock2.sock",
	"socket3": "/tmp/sock3.sock",
}

func main() {
	/*
		conn, err := net.Dial(protocol, sockAddr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		decoder := gob.NewDecoder(conn)
		encoder := gob.NewEncoder(conn)
	*/

	for {
		var text string
		fmt.Printf("input text > ")
		fmt.Scan(&text)
		if text == "exit" {
			break
		}
		sock := textsocket[text]
		conn, err := net.Dial(protocol, sock)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		decoder := gob.NewDecoder(conn)
		encoder := gob.NewEncoder(conn)

		m := &message.Echo{
			Length: len(text),
			Data:   []byte(text),
		}

		err = encoder.Encode(m)
		fmt.Println(m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[WRITE] ", m)

		err = decoder.Decode(m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[READ ] ", m)
	}
}
