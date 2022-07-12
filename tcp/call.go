package riptcp

import (
	"github.com/sharpvik/rip"
)

type call struct {
	*client
	function string
	argument interface{}
}

func (c *call) Arg(arg interface{}) rip.Call {
	c.argument = arg
	return c
}

func (c *call) Response() *rip.Response {
	req, e := rip.NewRequest(c.function, c.argument)
	if e != nil {
		return rip.ResponseError(e)
	}
	if e = c.Send(req); e != nil {
		return rip.ResponseError(e)
	}
	return ReadResponse(c.conn)
}

// Return checks if response contains an error, and if it does, returns
// that error straight away. Otherwise, it uses Unmarshal to decode response.
func (c *call) Return(v interface{}) rip.Error {
	resp := c.Response()
	if e := resp.Err(); e != nil {
		return e
	}
	return resp.Unmarshal(v)
}

// Err ignores the return value and reports response error if present.
func (c *call) Err() rip.Error {
	return c.Response().Err()
}
