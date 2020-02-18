package main

import (
	"fmt"
	"io"
	"net"
	"sso/data"
)

func main() {
	listener, err := net.Listen("tcp", data.ServerIP)
	if err != nil {
		fmt.Println("listern error!")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error!")
			continue
		}
		recvMessages(conn)
	}
}

func sendMessages() {

}

func recvMessages(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, "")
		if err != nil {
			fmt.Println("server recevice message error!")
		}
	}
}