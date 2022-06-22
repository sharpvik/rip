package riptcp

import (
	"github.com/sharpvik/rip/proto"
)

type Call struct {
	*Client
	function string
	argument interface{}
}

func (c *Call) Arg(arg interface{}) *Call {
	c.argument = arg
	return c
}

func (c *Call) Response() *proto.Response {
	req, e := proto.NewRequest(c.function, c.argument)
	if e != nil {
		return proto.ResponseError(e)
	}
	if e = c.Send(req); e != nil {
		return proto.ResponseError(e)
	}
	return ReadResponse(c.conn)
}

// Return checks if response contains an error, and if it does, returns
// that error straight away. Otherwise, it uses Unmarshal to decode response.
func (c *Call) Return(v interface{}) proto.Error {
	resp := c.Response()
	if e := resp.Err(); e != nil {
		return e
	}
	return resp.Unmarshal(v)
}

// Err ignores the return value and reports response error if present.
func (c *Call) Err() proto.Error {
	return c.Response().Err()
}
