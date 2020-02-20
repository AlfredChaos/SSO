package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"reflect"
)

func Register() {
	i := 0
	RegisterMap[i] = reflect.ValueOf(CheckSubSysLoginList)
	i++
	RegisterMap[i] = reflect.ValueOf(CheckGlobalLoginList)
	i++
	GlobalLoginUserList[ClientIP] = User1
}

func CheckGlobalLoginList(req *NotifyRequest) {
	err := checkGlobalLoginList(req.UserIP)
	if err != nil {
		fmt.Println("sso check error: ", err)
		return
	}
	return
}

func checkGlobalLoginList(userIP string) error {
	if userIP == "" {
		return errors.New("user ip is nil")
	}
	for k, _ := range GlobalLoginUserList {
		if k == userIP {
			fmt.Println("user has login")
			return nil
		}
	}
	return nil
}

func CheckSubSysLoginList(req *Request) {
	err := checkSubSysLoginList(req.IPAddr, req.AccIP)
	if err != nil {
		fmt.Println("server check error: ", err)
		return
	}
	return
}

func checkSubSysLoginList(userIP, accIP string) error {
	if userIP == "" {
		return errors.New("user ip is nil")
	}
	for k, _ := range SubSysLoginList {
		if k == userIP {
			fmt.Println("user has login")
			return nil
		}
	}
	err := redirect(userIP, accIP)
	if err != nil {
		return errors.New("server redirect error: " + fmt.Sprint(err))
	}
	return nil
}

func redirect(userIP, accIP string) error {
	conn, err := net.Dial("tcp", SSOIP)
	if err != nil {
		return errors.New("server dial sso error: " + fmt.Sprint(err))
	}
	defer conn.Close()

	rs, err := json.Marshal(&NotifyRequest{ReqID: 1, LocalIP: accIP, UserIP: userIP})
	if err != nil {
		return errors.New("server json Marshal error: " + fmt.Sprint(err))
	}
	_, err = conn.Write(rs)
	if err != nil {
		return errors.New("server redirect write data error: " + fmt.Sprint(err))
	}
	return nil
}