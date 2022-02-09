package rfip

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

type Request struct {
	FuncName string
	Argument []byte
}

func NewRequest(funcName string, argument interface{}) (r *Request, err error) {
	arg, err := json.Marshal(argument)
	if err != nil {
		return nil, ErrBadArgMarshal
	}
	r = &Request{
		FuncName: funcName,
		Argument: arg,
	}
	return
}

func (req *Request) String() string {
	bodySize := len(req.FuncName) + 1 + len(req.Argument)
	return fmt.Sprintf("%d %s %s", bodySize, req.FuncName, req.Argument)
}

func (req *Request) Bytes() []byte {
	return []byte(req.String())
}

func (req *Request) Send(w io.Writer) (n int, err error) {
	return w.Write(req.Bytes())
}

func ReadRequest(rd io.Reader) (req *Request, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = ErrUnexpectedPanic
		}
	}()

	bufr := bufio.NewReader(rd)

	contentLength, err := readContentLength(bufr)
	if err != nil {
		return
	}

	funcNameLength, funcName, err := readFuncName(bufr)
	if err != nil {
		return
	}

	argLength, err := calculateArgLength(contentLength, funcNameLength)
	if err != nil {
		return
	}

	arg, err := readBody(bufr, argLength)
	if err != nil {
		return
	}

	req = &Request{
		FuncName: funcName,
		Argument: arg,
	}
	return
}
