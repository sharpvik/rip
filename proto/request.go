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
	arg, err := json.Marshal(argument)
	if err != nil {
		return nil, WrapError(err, StatusBadRequest)
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

func (req *Request) Send(w io.Writer) (e Error) {
	if _, err := w.Write(req.Bytes()); err != nil {
		e = WrapError(err, StatusConnectionError)
	}
	return
}
