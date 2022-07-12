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

type Call interface {
	Arg(arg interface{}) Call
	Response() *Response
	Return(into interface{}) Error
	Err() Error
}
