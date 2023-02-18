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
	sockAddr = "/tmp/sock3.sock"
)

func main() {
	conn, err := net.Dial(protocol, sockAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	for {
		var text string
		fmt.Printf("input text > ")
		fmt.Scan(&text)
		if text == "exit" {
			break
		}

		m := &message.Echo{
			Length: len(text),
			Data:   []byte(text),
		}

		err = encoder.Encode(m)
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
