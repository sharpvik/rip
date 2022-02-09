package rfip

import (
	"bufio"
	"encoding/json"
	"io"
	"strconv"
)

type Response struct {
	Len  int
	Body []byte
}

func ResponseJSON(v interface{}) (resp *Response, err error) {
	body, err := json.Marshal(v)
	if err != nil {
		return
	}
	resp = &Response{
		Len:  len(body),
		Body: body,
	}
	return
}

func ResponseString(s string) *Response {
	return &Response{
		Len:  len(s),
		Body: []byte(s),
	}
}

func (resp *Response) Bytes() []byte {
	contentLength := []byte(strconv.Itoa(resp.Len) + " ")
	return append(contentLength, resp.Body...)
}

func (resp *Response) String() string {
	return string(resp.Bytes())
}

func (resp Response) Send(w io.Writer) (n int, err error) {
	return w.Write(resp.Bytes())
}

func ReadResponse(rd io.Reader) (resp *Response, err error) {
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

	body, err := readBody(bufr, contentLength)
	if err != nil {
		return
	}

	resp = &Response{
		Len:  contentLength,
		Body: body,
	}
	return
}
