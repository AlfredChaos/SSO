package server

import (
	"encoding/json"
	"fmt"
	"net"
	"sso/data"
)

func main() {
	listener, err := net.Listen("tcp", data.ServerIP)
	if err != nil {
		fmt.Println("listern error!", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error!", err)
			continue
		}
		recvMessages(conn)
	}
}

func sendMessages(ip string) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println("client dial server error!", err)
		return
	}
	defer conn.Close()
}

func recvMessages(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("server recev messages error!", err)
			return
		}
		user := &data.User{}
		err = json.Unmarshal(buf[:n], user)
		if err != nil {
			fmt.Println("server unmarshal error!", err)
		}
		fmt.Println(user.Username)
		fmt.Println(user.UserID)
	}
}