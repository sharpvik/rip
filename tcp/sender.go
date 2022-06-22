package riptcp

import (
	"io"

	"github.com/sharpvik/rip/proto"
)

func SendRequest(w io.Writer, req *proto.Request) (e proto.Error) {
	if _, err := w.Write(req.Bytes()); err != nil {
		e = proto.WrapError(err, proto.StatusConnectionError)
	}
	return
}

func SendResponse(w io.Writer, resp *proto.Response) (err error) {
	_, err = w.Write(resp.Bytes())
	return
}
