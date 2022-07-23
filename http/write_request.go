package riphttp

import (
	"bytes"
	"net/http"
	"path"

	"github.com/sharpvik/rip"
)

func WriteRequest(req *rip.Request) (r *http.Request, err error) {
	r, err = http.NewRequest(
		http.MethodGet,
		path.Join("/", req.Function),
		bytes.NewReader(req.Argument))
	return
}
