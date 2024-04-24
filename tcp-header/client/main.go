// 客户端端代码
package main

import (
	"encoding/binary"
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

	message1 := "Hello"
	message2 := "World"

	// 发送消息1
	messageBytes1 := []byte(message1)
	lengthBytes1 := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes1, uint32(len(messageBytes1)))

	_, err = conn.Write(lengthBytes1)
	if err != nil {
		fmt.Println("Error sending message length:", err.Error())
		return
	}
	fmt.Println("Sent message length:", len(messageBytes1))

	_, err = conn.Write(messageBytes1)
	if err != nil {
		fmt.Println("Error sending message:", err.Error())
		return
	}
	fmt.Println("Sent message:", message1)

	// 发送消息2
	messageBytes2 := []byte(message2)
	lengthBytes2 := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes2, uint32(len(messageBytes2)))

	_, err = conn.Write(lengthBytes2)
	if err != nil {
		fmt.Println("Error sending message length:", err.Error())
		return
	}
	fmt.Println("Sent message length:", len(messageBytes2))

	_, err = conn.Write(messageBytes2)
	if err != nil {
		fmt.Println("Error sending message:", err.Error())
		return
	}
	fmt.Println("Sent message:", message2)
}
