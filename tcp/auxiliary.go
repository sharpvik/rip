package riptcp

import (
	"io"

	"github.com/sharpvik/rip/proto"
)

func ReadRequest(rd io.Reader) (req *proto.Request, e proto.Error) {
	defer func() {
		if v := recover(); v != nil {
			e = proto.ErrUnexpectedPanic
		}
	}()
	return NewRequestReader(rd).ReadRequest()
}

func ReadResponse(rd io.Reader) (resp *proto.Response) {
	defer func() {
		if v := recover(); v != nil {
			resp = proto.ResponseError(proto.ErrUnexpectedPanic)
		}
	}()
	return NewReader(rd).ReadResponse()
}
