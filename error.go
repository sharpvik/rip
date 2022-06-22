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

// Error constants.
var (
	/* Bidirectional */

	ErrBadContentLengthRead = errors.New("bad body content length read")
	ErrInvalidContentLength = errors.New("invalid content length")
	ErrBadBodyRead          = errors.New("bad body read")

	/* Bad Request */

	ErrBadArgMarshal   = NewError("bad argument JSON marshal", StatusBadRequest)
	ErrBadArgUnmarshal = NewError("bad argument JSON unmarshal", StatusBadRequest)
	ErrBadFuncNameRead = NewError("err bad function name read", StatusBadRequest)
	ErrFuncNotFound    = NewError("function not found", StatusBadRequest)

	/* Bad Response */

	ErrBadResponseStatusRead = NewError("bad response status read", StatusBadResponse)

	/* Service Malfunction */

	ErrFuncWithBadArgc = NewError("function must have 0 or 1 argument", StatusServiceFailure)
	ErrFuncWithBadRetc = NewError("function must return 0 or 1 value", StatusServiceFailure)
	ErrBadBodyMarshal  = NewError("bad body JSON marshal", StatusServiceFailure)
	ErrUnexpectedPanic = NewError("unexpected panic occured", StatusServiceFailure)
)

func NewError(err string, status int) Error {
	return WrapError(errors.New(err), status)
}

func WrapError(err error, status int) Error {
	return &errorWithStatus{
		err:    err,
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
