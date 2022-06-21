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
	req, err := proto.NewRequest(c.function, c.argument)
	if err != nil {
		return proto.ResponseError(err)
	}
	if err = c.Send(req); err != nil {
		return proto.ResponseError(err)
	}
	return riptcp.ReadResponse(c.conn)
}

// Return checks if response contains an error, and if it does, returns
// that error straight away. Otherwise, it uses Unmarshal to decode response.
func (c *call) Return(v interface{}) proto.Error {
	resp := c.Response()
	if err := resp.Err(); err != nil {
		return err
	}
	return resp.Unmarshal(v)
}

// Err ignores the return value and reports response error if present.
func (c *call) Err() proto.Error {
	return c.Response().Err()
}
