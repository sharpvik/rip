package random

import (
	"github.com/sharpvik/rip"
	"github.com/sharpvik/rip/proto"
	riptcp "github.com/sharpvik/rip/tcp"
)

type Service struct {
	rip *riptcp.Client
}

func Client(addr string) *Service {
	return &Service{
		rip: riptcp.NewClient(addr),
	}
}

func (s *Service) RandInt() (i int, err proto.Error) {
	err = s.rip.Call("RandInt").Return(&i)
	return
}

func Server() *rip.Server {
	return rip.Use(Proto()).Server()
}
