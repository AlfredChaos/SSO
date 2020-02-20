package main

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"sso/data"
)

func main() {
	//初始化
	data.Register()

	//开始监听
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
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("server recev messages error: ", err)
		return
	}
	req := &data.Request{}
	err = json.Unmarshal(buf[:n], req)
	if err != nil {
		fmt.Println("server json unmarshal error: ", err)
		return
	}
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(req)
	data.RegisterMap[req.ReqID].Call(params)
}