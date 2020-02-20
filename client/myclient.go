package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sso/data"
)

func main() {
	accessServerRequest()

	// start listen, accept info
	/*listener, err := net.Listen("tcp", data.ServerIP)
	if err != nil {
		fmt.Println("listern error: ", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}

	}*/

}

func accessServerRequest() {
	conn, err := net.Dial("tcp", data.ServerIP)
	if err != nil {
		fmt.Println("client dial server error: ", err)
		return
	}
	defer conn.Close()

	rs, err := json.Marshal(&data.Request{ReqID: 0, IPAddr: data.ClientIP, AccIP: data.ServerIP})
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

// test code
func clientSendMessages(ip string) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println("client dial server error!", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client read data error!", err)
			return
		}
		rs, err := json.Marshal(line)
		if err != nil {
			fmt.Println("json Marshal error!")
		}
		_, _ = conn.Write(rs)
		if err != nil {
			fmt.Println("client wirte error!", err)
		}

	}
}