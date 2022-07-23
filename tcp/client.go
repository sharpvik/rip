package riptcp

import (
	"net"

	"github.com/sharpvik/rip"
)

type client struct {
	addr string
	conn net.Conn
}

func NewClient(addr string) rip.Client {
	return &client{
		addr: addr,
	}
}

func (c *client) Call(function string) rip.Call {
	return rip.UpgradeSimpleCall(&call{
		client:   c,
		function: function,
	})
}

func (c *client) Send(req *rip.Request) (e rip.Error) {
	if e = c.Connect(); e != nil {
		return
	}
	return SendRequest(c.conn, req)
}

func (c *client) Connect() (e rip.Error) {
	if c.conn == nil {
		e = c.Dial()
	}
	return
}

func (c *client) Dial() (e rip.Error) {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return rip.WrapError(err, rip.StatusConnectionError)
	}
	c.conn = conn
	return
}
