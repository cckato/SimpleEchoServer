package main
import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":12345"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Println("Start listening...")

	for  {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("Connected!!")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	for {
		messageBuf := make([]byte, 1024)
		messageLen, err := conn.Read(messageBuf)
		checkError(err)

		message := string(messageBuf[:messageLen])
		fmt.Println(message)
		message = message + " [from go server]"

		conn.Write([]byte(message))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}