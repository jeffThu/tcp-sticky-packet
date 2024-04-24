// 服务器端代码
package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		fmt.Println("Received message:", message)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on localhost:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		fmt.Println("New connection accepted.")

		go handleConnection(conn)
	}
}
