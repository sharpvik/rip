package riphttp

import (
	"io"
	"net/http"
	"strings"

	"github.com/sharpvik/rip"
)

func ReadRequest(r *http.Request) (req *rip.Request, e rip.Error) {
	argument, e := readArg(r)
	req = rip.NewRequestRaw(readFuncName(r), argument)
	return
}

func readFuncName(r *http.Request) string {
	return strings.Trim(r.URL.Path, "/")
}

func readArg(r *http.Request) (arg []byte, e rip.Error) {
	arg, err := io.ReadAll(r.Body)
	if err != nil {
		e = rip.WrapError(rip.ErrBadBodyRead, rip.StatusConnectionError)
	}
	return
}
