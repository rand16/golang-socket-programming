package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	protocol      = "unix"
	sockFixedAddr = "/tmp/command.sock"
)

func main() {
	printDoc()
	for {
		var command string
		var sockAddr string
		fmt.Printf("input command > ")
		fmt.Scan(&command)
		switch command {
		case "exit":
			fmt.Println("Bye")
			os.Exit(0)
		case "help":
			printDoc()
		default:
			sockAddr = initialize(command)
			//time.Sleep(10 * time.Second)
			fmt.Println(sockAddr)
			//commandExecution(sockAddr, command)
		}
	}
}

func printDoc() {
	fmt.Println("[hello]: 	Hello World")
	fmt.Println("[exit]: 	Exit this App")
	fmt.Println("[help]: 	Help Command")
}

func initialize(command string) string {
	conn, err := net.Dial(protocol, sockFixedAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	payload := `{"command":"` + command + `"}`
	request, err := http.NewRequest("GET", "http://localhost:8888", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Fatal(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		log.Fatal(err)
	}
	byteArray, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Get socket file > " + string(byteArray))

	return string(byteArray)
}

/*
func commandExecution(sockAddr string, command string) {
	switch command {
	case "hello":
		conn, err := net.Dial(protocol, sockAddr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		m := &message.Echo{
			Length: len(command),
			Data:   []byte(command),
		}
		err = m.Write(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[WRITE] ", m)

		err = m.Read(conn)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[READ ] ", m)
	}
}
*/
