package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"golang-socket-programming/pkg/message"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/command-prompt.sock"
)

func cleanup(socketFiles ...string) {
	for _, sock := range socketFiles {
		if _, err := os.Stat(sock); err == nil {
			if err := os.RemoveAll(sock); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	var socketFiles []string
	socketFiles = append(socketFiles, "/tmp/sock1.sock", "/tmp/sock2.sock", "/tmp/sock3.sock")
	cleanup(socketFiles...)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		fmt.Println("ctrl-c pressed!")
		close(quit)
		cleanup(socketFiles...)
		os.Exit(0)
	}()

	for _, file := range socketFiles {
		go initialize(file)
	}
	fmt.Scanln()
}

func initialize(file string) {
	listener, err := net.Listen(protocol, file)
	if err != nil {
		fmt.Println(file + " cannot be opened")
		log.Fatal(err)
	}
	fmt.Println("> Socket file launched: " + file)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(">>> accepted at %s: %s\n", file, conn.RemoteAddr().Network())
		switch file {
		case "/tmp/sock1.sock":
			go echo1(conn)
		case "/tmp/sock2.sock":
			go echo2(conn)
		case "/tmp/sock3.sock":
			go echo3(conn)
		}
	}
}

func echo1(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	fmt.Println("[MESSAGE] Call echo1 thread")

	for {
		m := &message.Echo{}
		err := decoder.Decode(m)
		if err != nil {
			if err == io.EOF {
				fmt.Println("=== closed by client")
				break
			}

			log.Println(err)
			break
		}

		fmt.Println("[READ ] ", m)

		s := strings.ToUpper(string(m.Data))
		m.Length = len(s)
		m.Data = []byte(s)

		err = encoder.Encode(m)
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("[WRITE] ", m)
	}
}

func echo2(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	fmt.Println("[MESSAGE] Call echo2 thread")

	for {
		m := &message.Echo{}
		err := decoder.Decode(m)
		if err != nil {
			if err == io.EOF {
				fmt.Println("=== closed by client")
				break
			}

			log.Println(err)
			break
		}

		fmt.Println("[READ ] ", m)

		s := strings.ToUpper(string(m.Data))
		m.Length = len(s)
		m.Data = []byte(s)

		err = encoder.Encode(m)
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("[WRITE] ", m)
	}
}
func echo3(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	fmt.Println("[MESSAGE] Call echo3 thread")

	for {
		m := &message.Echo{}
		err := decoder.Decode(m)
		if err != nil {
			if err == io.EOF {
				fmt.Println("=== closed by client")
				break
			}

			log.Println(err)
			break
		}

		fmt.Println("[READ ] ", m)

		s := strings.ToUpper(string(m.Data))
		m.Length = len(s)
		m.Data = []byte(s)

		err = encoder.Encode(m)
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("[WRITE] ", m)
	}
}
