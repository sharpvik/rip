package riphttp

import (
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/sharpvik/rip"
)

func ReadResponse(r *http.Response) *rip.Response {
	status, err := strconv.Atoi(Default(
		ripStatusOkString, r.Header.Get(ripStatusHeader)))
	if err != nil {
		return rip.ResponseError(rip.ErrBadResponseStatusRead)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return rip.ResponseError(rip.WrapError(
			rip.ErrBadBodyRead, rip.StatusConnectionError))
	}

	return &rip.Response{
		Status: status,
		Len:    len(body),
		Body:   body,
	}
}

func Default[T any](def, val T) T {
	if reflect.ValueOf(val).IsZero() {
		return def
	}
	return val
}
