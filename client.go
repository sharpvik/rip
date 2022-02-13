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

func (c *Client) MustInvoke(funcName string, arg interface{}) (resp *Response) {
	resp, err := c.Invoke(funcName, arg)
	PanicOnError(err)
	return
}

func (c *Client) Invoke(
	funcName string,
	arg interface{},
) (resp *Response, err error) {
	req, err := NewRequest(funcName, arg)
	if err != nil {
		return
	}
	if _, err = c.Send(req); err != nil {
		return
	}
	return ReadResponse(c.conn)
}

func (c *Client) Send(req *Request) (n int, err error) {
	if c.conn == nil {
		if err = c.Connect(); err != nil {
			return
		}
	}
	return req.Send(c.conn)
}

func (c *Client) Connect() (err error) {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		return
	}
	c.conn = conn
	return
}
