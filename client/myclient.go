package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sso/data"
)

func main() {
	clientSendMessages(data.ServerIP)
}

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