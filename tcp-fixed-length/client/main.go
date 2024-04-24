// 客户端端代码
package main

import (
	"fmt"
	"net"
)

const messageSize = 10

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server.")

	message1 := "Hello"
	message2 := "World"

	paddedMessage1 := fmt.Sprintf("%-10s", message1)
	paddedMessage2 := fmt.Sprintf("%-10s", message2)

	_, err = conn.Write([]byte(paddedMessage1))
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}
	fmt.Println("Sent message:", paddedMessage1)

	_, err = conn.Write([]byte(paddedMessage2))
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}
	fmt.Println("Sent message:", paddedMessage2)
}
