package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sso/data"
)

func main() {
	conn, err := net.Dial("tcp", data.ServerIP)
	if err != nil {
		fmt.Println("client dial server error!", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client read data error!", err)
			return
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("client wirte error!", err)
		}
	}
}

func clientSendMessages() {

}