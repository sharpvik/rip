package riptcp

import (
	"net"

	"github.com/sharpvik/rip"
)

type client struct {
	Addr string
	conn net.Conn
}

func NewClient(addr string) rip.Client {
	return &client{
		Addr: addr,
	}
}

func (c *client) Call(function string) rip.Call {
	return &call{
		client:   c,
		function: function,
	}
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
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		return rip.WrapError(err, rip.StatusConnectionError)
	}
	c.conn = conn
	return
}
