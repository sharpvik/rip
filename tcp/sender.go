package riptcp

import (
	"io"

	"github.com/sharpvik/rip"
)

func SendRequest(w io.Writer, req *rip.Request) (e rip.Error) {
	if _, err := w.Write(req.Bytes()); err != nil {
		e = rip.WrapError(err, rip.StatusConnectionError)
	}
	return
}

func SendResponse(w io.Writer, resp *rip.Response) (err error) {
	_, err = w.Write(resp.Bytes())
	return
}
