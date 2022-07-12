package riptcp

import (
	"github.com/sharpvik/rip"
)

func calculateArgLength(
	contentLength int,
	funcNameLength int,
) (length int, err error) {
	length = contentLength - funcNameLength
	if length < 0 {
		err = rip.ErrInvalidContentLength
	}
	return
}
