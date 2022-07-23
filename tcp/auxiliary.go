package riptcp

import (
	"io"

	"github.com/sharpvik/rip"
)

func ReadRequest(rd io.Reader) (req *rip.Request, e rip.Error) {
	defer func() {
		if v := recover(); v != nil {
			e = rip.ErrUnexpectedPanic
		}
	}()
	return NewRequestReader(rd).ReadRequest()
}

func ReadResponse(rd io.Reader) (resp *rip.Response) {
	defer func() {
		if v := recover(); v != nil {
			resp = rip.ResponseError(rip.ErrUnexpectedPanic)
		}
	}()
	return NewReader(rd).ReadResponse()
}

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
