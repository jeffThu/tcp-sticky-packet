// 服务器端代码
package main

import (
	"fmt"
	"net"
)

func handleConnection(connFd net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := connFd.Read(buffer) //协程做系统调用。堵塞读取文件描述符
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		data := buffer[:n]
		fmt.Println("Received data:", string(data))
	}
}

func main() {
	listenerFd, err := net.Listen("tcp", "localhost:8888") //进程向操作系统获取一个文件描述符
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listenerFd.Close()

	fmt.Println("Server started. Listening on localhost:8888")

	for {
		connFd, err := listenerFd.Accept() //进程向操作系统获取一个文件描述符
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		fmt.Println("New connection accepted.")

		go handleConnection(connFd) //一个文件描述符，一个协程。（这里可以优化为IO多路复用线程）
	}
}
