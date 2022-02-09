package rfip

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"path"
	"strings"
)

type Request struct {
	URL []string
	Arg []byte
}

func NewRequest(url string, argument interface{}) (r *Request, err error) {
	arg, err := json.Marshal(argument)
	if err != nil {
		return nil, ErrBadArgMarshal
	}
	r = &Request{
		URL: strings.Split(url, "/"),
		Arg: arg,
	}
	return
}

func (req *Request) String() string {
	url := path.Join(req.URL...)
	bodySize := len(url) + 1 + len(req.Arg)
	return fmt.Sprintf("%d %s %s", bodySize, url, req.Arg)
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

	urlLength, url, err := readURL(bufr)
	if err != nil {
		return
	}

	argLength, err := calculateArgLength(contentLength, urlLength)
	if err != nil {
		return
	}

	arg, err := readBody(bufr, argLength)
	if err != nil {
		return
	}

	req = &Request{
		URL: url,
		Arg: arg,
	}
	return
}
