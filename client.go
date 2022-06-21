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

func (c *Client) Send(req *proto.Request) (err proto.Error) {
	if err = c.Connect(); err != nil {
		return
	}
	return req.Send(c.conn)
}

func (c *Client) Connect() (err proto.Error) {
	if c.conn == nil {
		err = c.Dial()
	}
	return
}

func (c *Client) Dial() (err proto.Error) {
	conn, e := net.Dial("tcp", c.Addr)
	if e != nil {
		return proto.WrapError(e, proto.StatusConnectionError)
	}
	c.conn = conn
	return
}
