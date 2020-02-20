package data

type Request struct {
	ReqID  int
	IPAddr string
	AccIP  string
}

type NotifyRequest struct {
	ReqID   int
	LocalIP string
	UserIP  string
}