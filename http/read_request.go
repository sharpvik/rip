package riphttp

import (
	"io"
	"net/http"
	"strings"

	"github.com/sharpvik/rip/proto"
)

func ReadRequest(r *http.Request) (req *proto.Request, e proto.Error) {
	argument, e := readArg(r)
	req = proto.NewRequestRaw(readFuncName(r), argument)
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
