package rfip

import "errors"

var (
	ErrBadArgMarshal        = errors.New("bad argument JSON marshal")
	ErrBadContentLengthRead = errors.New("bad body content length read")
	ErrInvalidContentLength = errors.New("invalid content length")
	ErrBadURLRead           = errors.New("bad URL read")
	ErrBadBodyRead          = errors.New("bad body read")
	ErrUnexpectedPanic      = errors.New("unexpected panic occured")

	ErrFuncNotFound    = errors.New("function not found")
	ErrBadArgUnmarshal = errors.New("bad argument JSON unmarshal")
	ErrFuncWithBadArgc = errors.New("function must have 0 or 1 argument")
	ErrFuncWithBadRetc = errors.New("function must return 0 or 1 value")
)
