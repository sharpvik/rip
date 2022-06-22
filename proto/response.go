package proto

import (
	"bytes"
	"encoding/json"

	"github.com/sharpvik/rip/util"
)

type Response struct {
	Status int
	Len    int
	Body   []byte
}

func ResponseJSON(v interface{}) (resp *Response) {
	body, err := json.Marshal(v)
	if err != nil {
		return ResponseError(ErrBadBodyMarshal)
	}
	return ResponseOK(body)
}

func ResponseOK(body []byte) (resp *Response) {
	return ResponseRaw(StatusOK, body)
}

// ResponseError will panic on unknown status to encourage developers to use
// status constants declared in this package.
func ResponseError(err Error) (resp *Response) {
	return ResponseRaw(err.Status(), []byte(err.Error()))
}

func ResponseWrapError(err error, status int) (resp *Response) {
	return ResponseError(WrapError(err, status))
}

func ResponseRaw(status int, body []byte) (resp *Response) {
	return &Response{
		Status: status,
		Len:    len(body),
		Body:   body,
	}
}

func (resp *Response) OK() bool {
	return resp.Status == StatusOK
}

func (resp *Response) Bytes() []byte {
	var b bytes.Buffer
	b.Write(intWithSpaceAsBytes(resp.Status))
	b.Write(intWithSpaceAsBytes(resp.Len))
	b.Write(resp.Body)
	return b.Bytes()
}

func (resp *Response) String() string {
	return string(resp.Bytes())
}

func (resp *Response) Err() Error {
	if resp.OK() {
		return nil
	}
	return NewError(string(resp.Body), resp.Status)
}

func (resp *Response) MustUnmarshal(v interface{}) {
	util.PanicOnError(resp.Unmarshal(v))
}

func (resp *Response) Unmarshal(v interface{}) Error {
	if err := json.Unmarshal(resp.Body, v); err != nil {
		return WrapError(err, StatusServiceFailure)
	}
	return nil
}
