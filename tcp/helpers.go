package riptcp

import (
	"github.com/sharpvik/rip/proto"
)

func calculateArgLength(
	contentLength int,
	funcNameLength int,
) (length int, err proto.Error) {
	length = contentLength - funcNameLength
	if length < 0 {
		err = proto.ErrInvalidContentLength
	}
	return
}
