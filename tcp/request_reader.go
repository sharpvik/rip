package riptcp

import (
	"bufio"
	"io"

	"github.com/sharpvik/rip"
)

type RequestReader struct {
	*bufio.Reader
}

func NewRequestReader(rd io.Reader) (r *RequestReader) {
	return &RequestReader{bufio.NewReader(rd)}
}

func (r *RequestReader) ReadRequest() (req *rip.Request, e rip.Error) {
	contentLength, err := readInt(r.Reader)
	if err != nil {
		return nil, rip.WrapError(rip.ErrBadContentLengthRead, rip.StatusBadRequest)
	}

	funcNameLength, function, e := readFuncName(r.Reader)
	if e != nil {
		return
	}

	argLength, err := calculateArgLength(contentLength, funcNameLength)
	if e != nil {
		return nil, rip.WrapError(err, rip.StatusBadRequest)
	}

	argument, err := readBody(r.Reader, argLength)
	if e != nil {
		return nil, rip.WrapError(err, rip.StatusBadRequest)
	}

	return rip.NewRequestRaw(function, argument), nil
}
