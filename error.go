package rip

import (
	"errors"
)

type Error interface {
	Error() string
	Status() int
}

type errorWithStatus struct {
	err    error
	status int
}

func NewError(err string, status int) Error {
	return &errorWithStatus{
		err:    errors.New(err),
		status: status,
	}
}

func (ews *errorWithStatus) Error() string {
	return ews.err.Error()
}

func (ews *errorWithStatus) String() string {
	return ews.Error()
}

func (ews *errorWithStatus) Status() int {
	return ews.status
}

var (
	/* Generic errors */

	ErrBadArgMarshal = errors.New("bad argument JSON marshal")

	/* Response errors */

	ErrBadArgUnmarshal       = NewError("bad argument JSON unmarshal", StatusBadRequest)
	ErrBadContentLengthRead  = NewError("bad body content length read", StatusBadRequest)
	ErrBadResponseStatusRead = NewError("bad response status read", StatusBadRequest)
	ErrBadFuncNameRead       = NewError("err bad function name read", StatusBadRequest)
	ErrBadBodyRead           = NewError("bad body read", StatusBadRequest)
	ErrInvalidContentLength  = NewError("invalid content length", StatusBadRequest)
	ErrFuncNotFound          = NewError("function not found", StatusBadRequest)

	ErrFuncWithBadArgc = NewError("function must have 0 or 1 argument", StatusServiceMalfunction)
	ErrFuncWithBadRetc = NewError("function must return 0 or 1 value", StatusServiceMalfunction)
	ErrBadBodyMarshal  = NewError("bad body JSON marshal", StatusServiceMalfunction)
	ErrUnexpectedPanic = NewError("unexpected panic occured", StatusServiceMalfunction)
)
