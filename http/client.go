package riphttp

import (
	"net/http"

	"github.com/sharpvik/rip"
)

type client struct {
	addr string
}

func NewClient(addr string) rip.Client {
	return &client{
		addr: addr,
	}
}

func (c *client) Call(function string) rip.Call {
	return rip.UpgradeSimpleCall(&call{
		client:   c,
		function: function,
	})
}

func (c *client) Response(r *http.Request) (resp *rip.Response) {
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return rip.ResponseError(rip.WrapError(err, rip.StatusConnectionError))
	}
	return ReadResponse(res)
}
