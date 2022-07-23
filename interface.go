package rip

type Resolver interface {
	Handler
	Server() Server
}

type Handler interface {
	Handle(*Request) *Response
}

type Server interface {
	ListenAndServe(addr string) error
}

type Client interface {
	Call(function string) Call
}

type SimpleCall interface {
	SetArg(arg interface{})

	// Response attempts to invoke function on the remote server returning
	// response to the caller.
	Response() *Response
}

type Call interface {
	SimpleCall

	// Arg specifies a function call argument returning a reference to the same
	// Call.
	Arg(arg interface{}) Call

	// Return attempts to invoke the remote function, deserialising its return
	// value into the given reference. Error will be non-nil if something
	// went wrong.
	Return(into interface{}) Error

	// Err attempts to invoke the remote function discarding everything except
	// the error (or absence thereof) that it produces.
	Err() Error
}
