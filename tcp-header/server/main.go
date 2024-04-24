// 服务器端代码
package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	for {
		// 读取消息长度
		lengthBytes := make([]byte, 4)
		_, err := conn.Read(lengthBytes)
		if err != nil {
			fmt.Println("Error reading message length:", err.Error())
			return
		}

		// 解析消息长度
		length := binary.BigEndian.Uint32(lengthBytes)

		// 读取消息内容
		messageBytes := make([]byte, length)
		_, err = conn.Read(messageBytes)
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			return
		}

		message := string(messageBytes)
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
