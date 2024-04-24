// 客户端端代码
package main

import (
	"fmt"
	"net"
)

func main() {
	connFd, err := net.Dial("tcp", "localhost:8888") //进程向操作系统申请一个socket文件描述符
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer connFd.Close()

	fmt.Println("Connected to server.")

	//tcp连接建立完毕
	/*
		正常http协议通信：dns -> socket获取 -> 三次握手tcp连接建立 -> ssl通道建立 -> 从tcp连接中收发http应用层协议报文）

		* 但下文，我们的server与client间的应用层不是用http协议，而是自定义的应用层协议（自己定义收发格式、大小）
		* 自己定义应用层协议的话，就需要考虑几点：
			1.如果TCP粘包，怎么正确拆分
	*/

	message1 := "Hello"
	message2 := "World"

	_, err = connFd.Write([]byte(message1))
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}
	fmt.Println("Sent message:", message1)

	_, err = connFd.Write([]byte(message2))
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}
	fmt.Println("Sent message:", message2)
}
