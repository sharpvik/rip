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
