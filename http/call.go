package riphttp

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

	r, err := WriteRequest(req)
	if err != nil {
		return rip.ResponseError(rip.WrapError(err, rip.StatusBadRequest))
	}

	return c.client.Response(r)
}
