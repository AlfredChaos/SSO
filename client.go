package main

import (
	"fmt"
	"net"
	"sso/data"
)

func main() {
	conn, err := net.Dial("tcp", data.ServerIP)
	if err != nil {
		fmt.Println("client dial server error!")
	}
	
}

func clientSendMessages() {

}