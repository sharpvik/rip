package rip

import (
	"github.com/sharpvik/rip/proto"
	riptcp "github.com/sharpvik/rip/tcp"
)

type call struct {
	*Client
	function string
	argument interface{}
}

func (c *call) Arg(arg interface{}) *call {
	c.argument = arg
	return c
}

func (c *call) Response() *proto.Response {
	req, e := proto.NewRequest(c.function, c.argument)
	if e != nil {
		return proto.ResponseError(e)
	}
	if e = c.Send(req); e != nil {
		return proto.ResponseError(e)
	}
	return riptcp.ReadResponse(c.conn)
}

// Return checks if response contains an error, and if it does, returns
// that error straight away. Otherwise, it uses Unmarshal to decode response.
func (c *call) Return(v interface{}) proto.Error {
	resp := c.Response()
	if e := resp.Err(); e != nil {
		return e
	}
	return resp.Unmarshal(v)
}

// Err ignores the return value and reports response error if present.
func (c *call) Err() proto.Error {
	return c.Response().Err()
}
