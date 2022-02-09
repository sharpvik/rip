package rfip

import "errors"

var (
	ErrBadArgMarshal        = errors.New("bad argument JSON marshal")
	ErrBadContentLengthRead = errors.New("bad body content length read")
	ErrInvalidContentLength = errors.New("invalid content length")
	ErrBadURLRead           = errors.New("bad URL read")
	ErrBadBodyRead          = errors.New("bad body read")
	ErrUnexpectedPanic      = errors.New("unexpected panic occured")
)
