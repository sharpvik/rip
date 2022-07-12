package random

import (
	"github.com/sharpvik/rip"
	riptcp "github.com/sharpvik/rip/tcp"
)

type Client struct {
	rip *riptcp.Client
}

func NewClient(addr string) *Client {
	return &Client{
		rip: riptcp.NewClient(addr),
	}
}

func (s *Client) RandInt() (i int, err rip.Error) {
	err = s.rip.Call("RandInt").Return(&i)
	return
}
