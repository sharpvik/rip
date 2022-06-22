package rip

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Function string
	Argument []byte
}

func NewRequest(function string, arg interface{}) (*Request, Error) {
	argument, err := json.Marshal(arg)
	if err != nil {
		return nil, WrapError(err, StatusBadRequest)
	}
	return NewRequestRaw(function, argument), nil
}

func NewRequestRaw(function string, argument []byte) *Request {
	return &Request{
		Function: function,
		Argument: argument,
	}
}

func (req *Request) String() string {
	bodySize := len(req.Function) + 1 + len(req.Argument)
	return fmt.Sprintf("%d %s %s", bodySize, req.Function, req.Argument)
}

func (req *Request) Bytes() []byte {
	return []byte(req.String())
}
