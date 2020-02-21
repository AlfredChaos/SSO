package data

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"reflect"
	"time"
)

func Register() {
	RegisterMap[CheckSubSysLoginListIndex] = reflect.ValueOf(CheckSubSysLoginList)
	RegisterMap[CheckGlobalLoginListIndex] = reflect.ValueOf(CheckGlobalLoginList)
	RegisterMap[ClientLoginPageIndex] = reflect.ValueOf(ClientLoginPage)
	RegisterMap[CheckUserListIndex] = reflect.ValueOf(CheckUserList)
	RegisterMap[RedirectIndex] = reflect.ValueOf(Redirect)
	RegisterMap[CheckTokenIndex] = reflect.ValueOf(CheckToken)
	RegisterMap[UserLoginIndex] = reflect.ValueOf(UserLogin)
	RegisterMap[LoginAllowIndex] = reflect.ValueOf(LoginAllow)
}

func LoginAllow(req *Request) {
	loginAllow(req.Messages, req.Status)
	return
}

func UserLogin(req *Request) {
	err := userLogin(req.Cookies)
	if err != nil {
		fmt.Println("user login error: ", err)
		return
	}
	return
}

func CheckToken(req *Request) {
	err := checkToken(req.Token, req.LocalIP)
	if err != nil {
		fmt.Println("check token error: ", err)
		return
	}
	return
}

func Redirect(req *Request) {
	err := redirect("", ServerIP, req.Token, CheckTokenIndex)
	if err != nil {
		fmt.Println("server redirect to sso error: ", err)
		return
	}
	return
}

func CheckUserList(req *Request) {
	err := checkUserList(req.UserName, req.PassWord, req.LocalIP)
	if err != nil {
		fmt.Println("sso check User list error: ", err)
		return
	}
	return
}

func ClientLoginPage(req *Request) {
	err := clientLoginPage(req.AlreadyLogin, req.LocalIP)
	if err != nil {
		fmt.Println("client load login page error: ", err)
		return
	}
	return
}

func CheckGlobalLoginList(req *Request) {
	if req.LocalIP != "" {
		SSOCache.AccIP = req.LocalIP
	} else {
		fmt.Println("request access ip is nil")
		return
	}
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
	err := waitLogin(userIP)
	if err != nil {
		fmt.Println("sso waitLogin error: ", err)
		return nil
	}
	return nil
}

func CheckSubSysLoginList(req *Request) {
	err := checkSubSysLoginList(req.LocalIP, req.AccIP)
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
	err := redirect(userIP, accIP, "", CheckGlobalLoginListIndex)
	if err != nil {
		return errors.New("server redirect error: " + fmt.Sprint(err))
	}
	return nil
}

func redirect(userIP, accIP, token string, requestID int) error {
	conn, err := net.Dial("tcp", SSOIP)
	if err != nil {
		return errors.New("server dial sso error: " + fmt.Sprint(err))
	}
	defer conn.Close()

	rs, err := json.Marshal(&Request{ReqID: requestID, LocalIP: accIP, UserIP: userIP, Token: token})
	if err != nil {
		return errors.New("server json Marshal error: " + fmt.Sprint(err))
	}
	_, err = conn.Write(rs)
	if err != nil {
		return errors.New("server redirect write data error: " + fmt.Sprint(err))
	}
	return nil
}

func waitLogin(userIP string) error {
	conn, err := net.Dial("tcp", userIP)
	if err != nil {
		return errors.New("sso dial client error: " + fmt.Sprint(err))
	}
	defer conn.Close()

	rs, err := json.Marshal(&Request{ReqID: 2, LocalIP: SSOIP, AlreadyLogin: false})
	if err != nil {
		return errors.New("sso json Marshal error: " + fmt.Sprint(err))
	}
	_, err = conn.Write(rs)
	if err != nil {
		return errors.New("sso waitLogin write data error: " + fmt.Sprint(err))
	}
	return nil
}

func clientLoginPage(alreadyLogin bool, ssoIP string) error {
	if !alreadyLogin {
		var user, password string
		fmt.Print("user: ")
		_, _ = fmt.Scanf("%s", &user)
		fmt.Print("passwd: ")
		_, _ = fmt.Scanf("%s", &password)

		conn, err := net.Dial("tcp", ssoIP)
		if err != nil {
			return errors.New("client dial sso error: " + fmt.Sprint(err))
		}
		defer conn.Close()

		rs, err := json.Marshal(&Request{ReqID: 3, UserName: user, PassWord: password, LocalIP: ClientIP})
		if err != nil {
			return errors.New("client login info marshal error: " + fmt.Sprint(err))
		}
		_, err = conn.Write(rs)
		if err != nil {
			return errors.New("client login write data error: " + fmt.Sprint(err))
		}
		return nil
	} else {
		fmt.Println("You has login!")
		return nil
	}
}

func checkUserList(username, password, clientIP string) error {
	for _, user := range UserList {
		if user.Username == username {
			if user.Password == password {
				token, err := createGlobalSession(clientIP)
				if err != nil {
					return errors.New("generate token error: " + fmt.Sprint(err))
				}
				err = callServer(token)
				if err != nil {
					return errors.New("sso call server error: " + fmt.Sprint(err))
				}
			} else {
				return errors.New("password error")
			}
		} else {
			return errors.New("user not exist")
		}
	}
	return nil
}

func createGlobalSession(userIP string) (string, error) {
	cookie := &Session{
		SessionID: 1,
		SrcIPAddr: userIP,
		DstIPAddr: SSOCache.AccIP,
		SsoIPAddr: SSOIP,
		Status:    KeepAlive,
	}
	GlobalCookies[0] = cookie
	//使用md5加密假装生成token
	w := md5.New()
	_, err := io.WriteString(w, fmt.Sprintf("%s%s%s%s%s", string(cookie.SessionID), cookie.SrcIPAddr, cookie.DstIPAddr, cookie.SsoIPAddr, string(time.Now().Day())))
	if err != nil {
		return "", errors.New("sso encrypted error: " + fmt.Sprint(err))
	}
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str, nil
}

func callServer(token string) error {
	conn, err := net.Dial("tcp", ServerIP)
	if err != nil {
		return errors.New("sso dial server error: " + fmt.Sprint(err))
	}
	defer conn.Close()

	rs, err := json.Marshal(&Request{ReqID: 4, Token: token})
	if err != nil {
		return errors.New("sso json Marshal error: " + fmt.Sprint(err))
	}
	_, err = conn.Write(rs)
	if err != nil {
		return errors.New("sso callServer write data error: " + fmt.Sprint(err))
	}
	return nil
}

func checkToken(token, serverIP string) error {
	//验证token
	cookie := GlobalCookies[0]
	w := md5.New()
	_, err := io.WriteString(w, fmt.Sprintf("%s%s%s%s%s", string(cookie.SessionID), cookie.SrcIPAddr, serverIP, cookie.SsoIPAddr, string(time.Now().Day())))
	if err != nil {
		return errors.New("sso encrypted error: " + fmt.Sprint(err))
	}
	checksum := fmt.Sprintf("%x", w.Sum(nil))
	fmt.Println(checksum)
	fmt.Println(token)

	if checksum == token {
		User1.Status = HasLogin
		GlobalLoginUserList[cookie.SrcIPAddr] = User1
	} else {
		cookie.Status = Destroy
	}

	//让server知道token有效or无效
	conn, err := net.Dial("tcp", ServerIP)
	if err != nil {
		return errors.New("sso call server err: " + fmt.Sprint(err))
	}
	defer conn.Close()

	rs, err := json.Marshal(&Request{ReqID: UserLoginIndex, Cookies: cookie})
	if err != nil {
		return errors.New("sso marshal token error: " + fmt.Sprint(err))
	}
	_, err = conn.Write(rs)
	if err != nil {
		return errors.New("sso checkToken write data error: " + fmt.Sprint(err))
	}
	return nil
}

func userLogin(cookie *Session) error {
	if cookie == nil {
		return errors.New("server get cookie nil")
	}
	req := &Request{ReqID: LoginAllowIndex}
	if cookie.Status == KeepAlive {
		req.Status = true
		req.Messages = "login success"
		PrivateCookies[0] = cookie
		SubSysLoginList[cookie.SrcIPAddr] = User1
	} else if cookie.Status == Destroy {
		req.Status = false
		req.Messages = "Login refused"
	}

	conn, err := net.Dial("tcp", cookie.SrcIPAddr)
	if err != nil {
		return errors.New("server dial client error: " + fmt.Sprint(err))
	}
	defer conn.Close()
	rs, err := json.Marshal(req)
	if err != nil {
		return errors.New("server marshal request error: " + fmt.Sprint(err))
	}
	_, err = conn.Write(rs)
	if err != nil {
		return errors.New("server userLogin write data error: " + fmt.Sprint(err))
	}
	return nil
}

func loginAllow(message string, status bool) {
	fmt.Println(message)

	if !status {
		os.Exit(1)
	}
}