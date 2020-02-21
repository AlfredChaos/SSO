package data

type Request struct {
	ReqID        int
	LocalIP      string
	AccIP        string
	UserName     string
	PassWord     string
	AlreadyLogin bool
	UserIP       string
	Token        string
	Cookies      *Session
	Messages     string
	Status       bool
}

/*type NotifyRequest struct {
	ReqID        int
	LocalIP      string
	UserIP       string
	AlreadyLogin bool
}*/