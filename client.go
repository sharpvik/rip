package rip

import (
	"net"

	"github.com/sharpvik/rip/proto"
)

type Client struct {
	Addr string
	conn net.Conn
}

func NewClient(addr string) *Client {
	return &Client{
		Addr: addr,
	}
}

func (c *Client) Call(function string) *call {
	return &call{
		Client:   c,
		function: function,
	}
}

func (c *Client) Send(req *proto.Request) (e proto.Error) {
	if e = c.Connect(); e != nil {
		return
	}
	return req.Send(c.conn)
}

func (c *Client) Connect() (e proto.Error) {
	if c.conn == nil {
		e = c.Dial()
	}
	return
}

func (c *Client) Dial() (e proto.Error) {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		return proto.WrapError(err, proto.StatusConnectionError)
	}
	c.conn = conn
	return
}
