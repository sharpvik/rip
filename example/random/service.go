package random

import (
	"github.com/sharpvik/rip"
)

type service struct {
	rip *rip.Client
}

func Service(addr string) *service {
	return &service{
		rip: rip.NewClient(addr),
	}
}

func (s *service) RandInt() (i int, err rip.Error) {
	err = s.rip.MustInvoke("RandInt", nil).Return(&i)
	return
}

func Server() *rip.Server {
	return rip.Use(Proto()).Server()
}
