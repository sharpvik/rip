package riptcp

import (
	"net"

	"github.com/sharpvik/rip"
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

func (c *Client) Call(function string) *Call {
	return &Call{
		Client:   c,
		function: function,
	}
}

func (c *Client) Send(req *rip.Request) (e rip.Error) {
	if e = c.Connect(); e != nil {
		return
	}
	return SendRequest(c.conn, req)
}

func (c *Client) Connect() (e rip.Error) {
	if c.conn == nil {
		e = c.Dial()
	}
	return
}

func (c *Client) Dial() (e rip.Error) {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		return rip.WrapError(err, rip.StatusConnectionError)
	}
	c.conn = conn
	return
}
