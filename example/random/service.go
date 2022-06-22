package random

import (
	"github.com/sharpvik/rip"
	"github.com/sharpvik/rip/proto"
)

type Service struct {
	rip *rip.Client
}

func Client(addr string) *Service {
	return &Service{
		rip: rip.NewClient(addr),
	}
}

func (s *Service) RandInt() (i int, err proto.Error) {
	err = s.rip.Call("RandInt").Return(&i)
	return
}

func Server() *rip.Server {
	return rip.Use(Proto()).Server()
}
