package rip

import (
	"net"
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

func (c *Client) Send(req *Request) (err Error) {
	if err = c.Connect(); err != nil {
		return
	}
	return req.Send(c.conn)
}

func (c *Client) Connect() (err Error) {
	if c.conn == nil {
		err = c.Dial()
	}
	return
}

func (c *Client) Dial() (err Error) {
	conn, e := net.Dial("tcp", c.Addr)
	if e != nil {
		return WrapError(e, StatusConnectionError)
	}
	c.conn = conn
	return
}
