package random

import (
	"github.com/sharpvik/rip"
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

func (s *Service) RandInt() (i int, err rip.Error) {
	err = s.rip.Call("RandInt").Return(&i)
	return
}
