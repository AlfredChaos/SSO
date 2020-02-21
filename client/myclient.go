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
	clientInit()

	accessServerRequest()

	// start listen, accept info
	listener, err := net.Listen("tcp", data.ClientIP)
	if err != nil {
		fmt.Println("listener error: ", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}
		clientRecvMessages(conn)
	}

}

func clientInit() {
	data.Register()
}

func accessServerRequest() {
	conn, err := net.Dial("tcp", data.ServerIP)
	if err != nil {
		fmt.Println("client dial server error: ", err)
		return
	}
	defer conn.Close()

	rs, err := json.Marshal(&data.Request{ReqID: 0, LocalIP: data.ClientIP, AccIP: data.ServerIP})
	if err != nil {
		fmt.Println("client access server error: ", err)
		return
	}
	_, err = conn.Write(rs)
	if err != nil {
		fmt.Println("client wirte error: ", err)
		return
	}
}

func clientRecvMessages(conn net.Conn) {
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