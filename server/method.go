package server

import (
	"errors"
	"sso/data"
)

func checkLoginList(userIP string) error {
	for k, _ := range data.LoginUserList {
		if k == userIP {
			return errors.New("user has login")
		}
	}
	redirect(userIP)
}

func redirect(userIP string) {

}