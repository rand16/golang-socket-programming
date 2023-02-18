package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

const (
	protocol      = "unix"
	sockFixedAddr = "/tmp/command.sock"
)

// はじめにどのソケットファイルを指定するかを決定するために，コマンドを読み取り対応するソケットファイルを返す操作を実行する
type Init struct {
	Command string
}

// コマンド名とソケットファイルの対応関係
var CommandSockPath = map[string]string{
	"hello": "/tmp/hello.sock",
}

func cleanup(socket ...string) {
	for _, sock := range socket {
		if _, err := os.Stat(sock); err == nil {
			if err := os.RemoveAll(sock); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	cleanup(sockFixedAddr)

	// ctrl-c シグナル受信
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		fmt.Println("ctrl-c pressed!")
		close(quit)
		cleanup(sockFixedAddr)
		os.Exit(0)
	}()

	// 2023/02/17 Initialize関数でまとめてみる
	sockAddr := initialize()
	fmt.Println("Resolve socket file >  " + sockAddr)
	// ↑sockAddrをチャネル送信にしてみる．go echo(conn)の形に戻る
	// <- 失敗
}

func initialize() string {
	commandListener, err := net.Listen(protocol, sockFixedAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer commandListener.Close()

	fmt.Println("> Server launched")
	conn, err := commandListener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	sockAddr := resolveSocketFile(conn)

	return sockAddr
}

func resolveSocketFile(conn net.Conn) string {
	defer conn.Close()

	fmt.Println(">>> accepted: ", conn.RemoteAddr().Network())
	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Fatal(err)
	}
	byteArray, _ := ioutil.ReadAll(request.Body)
	output := &Init{}
	if err = json.Unmarshal(byteArray, output); err != nil {
		log.Fatal(err)
	}

	sockAddr := generateSockPath(output.Command)

	response := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       ioutil.NopCloser(strings.NewReader(sockAddr + "\n")),
	}
	response.Write(conn)
	conn.Close()

	fmt.Println("<<< finish Initialization")

	return sockAddr
}

func generateSockPath(command string) string {
	sockAddr := "/tmp/" + command + ".sock"
	return sockAddr
}

/*
func basic(sockAddr string) {
	basicListener, err := net.Listen(protocol, sockAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer basicListener.Close()

	fmt.Println("> basic thread")
	conn, err := basicListener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		m := &message.Echo{}
		err := m.Read(conn)
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

		err = m.Write(conn)
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("[WRITE] ", m)
	}
}
*/
