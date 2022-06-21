package riptcp

import (
	"bufio"
	"io"

	"github.com/sharpvik/rip/proto"
)

type RequestReader struct {
	*bufio.Reader
}

func NewRequestReader(rd io.Reader) (r *RequestReader) {
	return &RequestReader{bufio.NewReader(rd)}
}

func (r *RequestReader) ReadRequest() (req *proto.Request, e proto.Error) {
	contentLength, err := readInt(r.Reader)
	if err != nil {
		return nil, proto.WrapError(proto.ErrBadContentLengthRead, proto.StatusBadRequest)
	}

	funcNameLength, funcName, e := readFuncName(r.Reader)
	if e != nil {
		return
	}

	argLength, err := calculateArgLength(contentLength, funcNameLength)
	if e != nil {
		return nil, proto.WrapError(err, proto.StatusBadRequest)
	}

	arg, err := readBody(r.Reader, argLength)
	if e != nil {
		return nil, proto.WrapError(err, proto.StatusBadRequest)
	}

	return &proto.Request{
		Function: funcName,
		Argument: arg,
	}, nil
}
