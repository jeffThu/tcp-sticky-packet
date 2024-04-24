// 客户端端代码
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server.")

	message1 := "Hello\n"
	message2 := "World\n"

	_, err = conn.Write([]byte(message1))
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}
	fmt.Println("Sent message:", message1)

	_, err = conn.Write([]byte(message2))
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}
	fmt.Println("Sent message:", message2)
}
