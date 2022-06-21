package rip

import (
	"github.com/sharpvik/rip/proto"
)

type Resolver interface {
	Handle(*proto.Request) *proto.Response
	Server() *Server
}

type resolver struct {
	Handler
}

func NewResolver(h Handler) Resolver {
	return &resolver{Handler: h}
}

func Use(source interface{}) Resolver {
	return NewResolver(NewHanlder(source))
}

func (r resolver) Server() *Server {
	return NewServer(r)
}
