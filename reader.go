package rip

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Reader struct {
	*bufio.Reader
}

func NewReader(rd io.Reader) (r *Reader) {
	return &Reader{bufio.NewReader(rd)}
}

func (r *Reader) ReadRequest() (req *Request, err Error) {
	contentLength, err := r.readContentLength()
	if err != nil {
		return
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

	return &Request{
		FuncName: funcName,
		Argument: arg,
	}, nil
}

func (r *Reader) ReadResponse() (resp *Response, err error) {
	responseStatus, err := r.readResponseStatus()
	if err != nil {
		return
	}

	contentLength, err := r.readContentLength()
	if err != nil {
		return
	}

	body, err := r.readBody(contentLength)
	if err != nil {
		return
	}

	return &Response{
		Status: responseStatus,
		Len:    contentLength,
		Body:   body,
	}, nil
}

func (r *Reader) readResponseStatus() (status int, err error) {
	status, err = r.readInt()
	if err != nil {
		err = ErrBadResponseStatusRead
	}
	return
}

func (r *Reader) readContentLength() (contentLength int, err Error) {
	contentLength, e := r.readInt()
	if e != nil {
		err = ErrBadContentLengthRead
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

func (r *Reader) readFuncName() (length int, name string, err Error) {
	withSpace, e := r.ReadString(' ')
	if e != nil {
		err = ErrBadFuncNameRead
		return
	}
	length = len(withSpace)
	name = strings.TrimSpace(withSpace)
	return
}

func (r *Reader) readBody(contentLength int) (data []byte, err Error) {
	data = make([]byte, contentLength)
	if _, e := io.ReadFull(r, data); e != nil {
		err = ErrBadBodyRead
	}
	return
}
