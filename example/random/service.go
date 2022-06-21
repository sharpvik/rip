package random

import (
	"github.com/sharpvik/rip"
	"github.com/sharpvik/rip/proto"
)

type service struct {
	rip *rip.Client
}

func Service(addr string) *service {
	return &service{
		rip: rip.NewClient(addr),
	}
}

func (s *service) RandInt() (i int, err proto.Error) {
	err = s.rip.Call("RandInt").Return(&i)
	return
}

func Server() *rip.Server {
	return rip.Use(Proto()).Server()
}
