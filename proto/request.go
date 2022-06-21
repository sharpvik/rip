package proto

import (
	"encoding/json"
	"fmt"
	"io"
)

type Request struct {
	Function string
	Argument []byte
}

func NewRequest(function string, argument interface{}) (*Request, Error) {
	arg, e := json.Marshal(argument)
	if e != nil {
		return nil, WrapError(e, StatusBadRequest)
	}
	return &Request{
		Function: function,
		Argument: arg,
	}, nil
}

func (req *Request) String() string {
	bodySize := len(req.Function) + 1 + len(req.Argument)
	return fmt.Sprintf("%d %s %s", bodySize, req.Function, req.Argument)
}

func (req *Request) Bytes() []byte {
	return []byte(req.String())
}

func (req *Request) Send(w io.Writer) (err Error) {
	if _, e := w.Write(req.Bytes()); e != nil {
		return WrapError(e, StatusConnectionError)
	}
	return
}
