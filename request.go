package rfip

import (
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

func ReadRequest(rd io.Reader) (req *Request, err Error) {
	defer func() {
		if v := recover(); v != nil {
			err = ErrUnexpectedPanic
		}
	}()
	return NewReader(rd).ReadRequest()
}
