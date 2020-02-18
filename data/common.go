package data

type Login struct {
	Username string
	Password string
	IpAddr   string
	Status	 int
}

const (
	NoLogin = iota
	HasLogin
	Logout
)

type User struct {
	UserID   int
	Username string
	Password string
	Status   int
}

type Session struct {
	SessionID int
	SrcIPAddr string
	DstIPAddr string
	Status    int
}

var UserList []*User

var LoginUserList []*User

var User1 *User = &User{
	UserID:   1,
	Username: "cao",
	Password: "cao",
	Status:   0,
}

const (
	KeepAlive = iota
	Destroy
)

var ServerIP string = "127.0.0.1:9000"
var SSOIP string = "127.0.0.1:9001"