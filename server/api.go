package server

type Service struct {}

func (s *Service) CheckLoginList() {
	err := checkLoginList()
}
