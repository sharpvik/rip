package riptcp

import (
	"github.com/sharpvik/rip"
)

type call struct {
	*client
	function string
	argument interface{}
}

func (c *call) SetArg(arg interface{}) {
	c.argument = arg
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
