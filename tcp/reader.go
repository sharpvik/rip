package riptcp

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/sharpvik/rip/proto"
)

type Reader struct {
	*bufio.Reader
}

func ReadRequest(rd io.Reader) (req *proto.Request, err proto.Error) {
	defer func() {
		if v := recover(); v != nil {
			err = proto.ErrUnexpectedPanic
		}
	}()
	return NewReader(rd).ReadRequest()
}

func ReadResponse(rd io.Reader) (resp *proto.Response) {
	defer func() {
		if v := recover(); v != nil {
			resp = proto.ResponseError(proto.ErrUnexpectedPanic)
		}
	}()
	return NewReader(rd).ReadResponse()
}

func NewReader(rd io.Reader) (r *Reader) {
	return &Reader{bufio.NewReader(rd)}
}

func (r *Reader) ReadRequest() (req *proto.Request, err proto.Error) {
	contentLength, e := r.readInt()
	if e != nil {
		return nil, proto.ErrBadContentLengthRead
	}

	funcNameLength, funcName, err := r.readFuncName()
	if err != nil {
		return
	}

	argLength, err := calculateArgLength(contentLength, funcNameLength)
	if err != nil {
		return
	}

	arg, err := r.readBody(argLength)
	if err != nil {
		return
	}

	return &proto.Request{
		Function: funcName,
		Argument: arg,
	}, nil
}

func (r *Reader) ReadResponse() (resp *proto.Response) {
	responseStatus, err := r.readResponseStatus()
	if err != nil {
		return proto.ResponseError(err)
	}

	contentLength, e := r.readInt()
	if e != nil {
		return proto.ResponseError(proto.WrapError(
			proto.ErrBadContentLengthRead, proto.StatusBadResponse))
	}

	body, err := r.readBody(contentLength)
	if err != nil {
		return proto.ResponseError(err)
	}

	return &proto.Response{
		Status: responseStatus,
		Len:    contentLength,
		Body:   body,
	}
}

func (r *Reader) readResponseStatus() (status int, err proto.Error) {
	status, e := r.readInt()
	if e != nil {
		err = proto.WrapError(proto.ErrBadResponseStatusRead, proto.StatusBadResponse)
	}
	return
}

func (r *Reader) readInt() (i int, err error) {
	str, err := r.ReadString(' ')
	if err != nil {
		return
	}
	return strconv.Atoi(strings.TrimSpace(str))
}

func (r *Reader) readFuncName() (length int, name string, err proto.Error) {
	withSpace, e := r.ReadString(' ')
	if e != nil {
		err = proto.ErrBadFuncNameRead
		return
	}
	length = len(withSpace)
	name = strings.TrimSpace(withSpace)
	return
}

func (r *Reader) readBody(contentLength int) (data []byte, err proto.Error) {
	data = make([]byte, contentLength)
	if _, e := io.ReadFull(r, data); e != nil {
		err = proto.ErrBadBodyRead
	}
	return
}
