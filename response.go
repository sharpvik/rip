package rip

import (
	"encoding/json"
	"io"
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
	return ResponseWithStatus(StatusOK, body)
}

// ResponseError will panic on unknown status to encourage developers to use
// status constants declared in this package.
func ResponseError(err Error) (resp *Response) {
	return ResponseWithStatus(err.Status(), []byte(err.Error()))
}

func ResponseWithStatus(status int, body []byte) (resp *Response) {
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
	responseStatus := intWithSpaceAsBytes(resp.Status)
	contentLength := intWithSpaceAsBytes(resp.Len)
	metadata := append(responseStatus, contentLength...)
	return append(metadata, resp.Body...)
}

func (resp *Response) String() string {
	return string(resp.Bytes())
}

func (resp *Response) Err() Error {
	if resp.Status == StatusOK {
		return nil
	}
	return NewError(string(resp.Body), resp.Status)
}

func (resp *Response) Send(w io.Writer) (err error) {
	_, err = w.Write(resp.Bytes())
	return
}

func (resp *Response) MustUnmarshal(v interface{}) {
	PanicOnError(resp.Unmarshal(v))
}

// Return checks if response contains an error, and if it does, returns
// that error straight away. Otherwise, it uses Response.Unmarshal to decode
// response body.
func (resp *Response) Return(v interface{}) Error {
	if err := resp.Err(); err != nil {
		return err
	}
	return resp.Unmarshal(v)
}

func (resp *Response) Unmarshal(v interface{}) Error {
	if err := json.Unmarshal(resp.Body, v); err != nil {
		return NewError(err.Error(), StatusServiceMalfunction)
	}
	return nil
}

func ReadResponse(rd io.Reader) (resp *Response, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = ErrUnexpectedPanic
		}
	}()
	return NewReader(rd).ReadResponse()
}
