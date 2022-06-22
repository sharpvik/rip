package riptcp

import (
	"github.com/sharpvik/rip"
)

type Resolver interface {
	rip.Handler
	Server() *Server
}

type resolver struct {
	rip.Handler
}

func NewResolver(h rip.Handler) Resolver {
	return &resolver{Handler: h}
}

func Use(source interface{}) Resolver {
	return NewResolver(rip.NewHanlder(source))
}

func (r resolver) Server() *Server {
	return NewServer(r)
}
