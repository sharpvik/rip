package riphttp

import (
	"io"
	"net/http"
	"strings"

	"github.com/sharpvik/rip/proto"
)

type RequestReader struct {
	req *http.Request
}

func NewRequestReader(r *http.Request) *RequestReader {
	return &RequestReader{
		req: r,
	}
}

func (r *RequestReader) ReadRequest() (req *proto.Request, e proto.Error) {
	argument, e := readArg(r.req)
	req = proto.NewRequestRaw(readFuncName(r.req), argument)
	return
}

func readFuncName(r *http.Request) string {
	return strings.Trim(r.URL.Path, "/")
}

func readArg(r *http.Request) (arg []byte, e proto.Error) {
	arg, err := io.ReadAll(r.Body)
	if err != nil {
		e = proto.WrapError(proto.ErrBadBodyRead, proto.StatusConnectionError)
	}
	return
}
