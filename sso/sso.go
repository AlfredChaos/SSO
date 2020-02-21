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
	ssoInit()

	//开始监听
	listener, err := net.Listen("tcp", data.SSOIP)
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
		ssoRecvMessages(conn)
	}
}

func ssoInit() {
	data.Register()
	data.UserList[0] = data.User1
	data.SSOCache = &data.Request{}
}

func ssoRecvMessages(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("server recev messages error: ", err)
		return
	}
	notify := &data.Request{}
	err = json.Unmarshal(buf[:n], notify)
	if err != nil {
		fmt.Println("sso json unmarshl request error: ", err)
		return
	}
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(notify)
	data.RegisterMap[notify.ReqID].Call(params)
}