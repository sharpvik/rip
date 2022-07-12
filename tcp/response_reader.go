package riptcp

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/sharpvik/rip"
)

type ResponseReader struct {
	*bufio.Reader
}

func NewReader(rd io.Reader) (r *ResponseReader) {
	return &ResponseReader{bufio.NewReader(rd)}
}

func (r *ResponseReader) ReadResponse() (resp *rip.Response) {
	responseStatus, e := r.readResponseStatus()
	if e != nil {
		return rip.ResponseError(e)
	}

	contentLength, err := readInt(r.Reader)
	if err != nil {
		return rip.ResponseWrapError(
			rip.ErrBadContentLengthRead, rip.StatusBadResponse)
	}

	body, err := readBody(r.Reader, contentLength)
	if e != nil {
		return rip.ResponseWrapError(err, rip.StatusBadResponse)
	}

	return &rip.Response{
		Status: responseStatus,
		Len:    contentLength,
		Body:   body,
	}
}

func (r *ResponseReader) readResponseStatus() (status int, e rip.Error) {
	status, err := readInt(r.Reader)
	if err != nil {
		e = rip.ErrBadResponseStatusRead
	}
	return
}

func readInt(r *bufio.Reader) (i int, err error) {
	str, err := r.ReadString(' ')
	if err != nil {
		return
	}
	return strconv.Atoi(strings.TrimSpace(str))
}

func readFuncName(r *bufio.Reader) (length int, name string, e rip.Error) {
	withSpace, err := r.ReadString(' ')
	if err != nil {
		e = rip.ErrBadFuncNameRead
		return
	}
	length = len(withSpace)
	name = strings.TrimSpace(withSpace)
	return
}

func readBody(r *bufio.Reader, contentLength int) (data []byte, err error) {
	data = make([]byte, contentLength)
	if _, err = io.ReadFull(r, data); err != nil {
		err = rip.ErrBadBodyRead
	}
	return
}
