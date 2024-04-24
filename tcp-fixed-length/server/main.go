// 服务器端代码
package main

import (
	"fmt"
	"net"
)

const messageSize = 10

func handleConnection(conn net.Conn) {
	buffer := make([]byte, messageSize)

	for {
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		message := string(buffer)
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
