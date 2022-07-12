package riptcp

import (
	"github.com/sharpvik/rip"
)

type resolver struct {
	rip.Handler
}

func NewResolver(h rip.Handler) rip.Resolver {
	return &resolver{Handler: h}
}

func Use(source interface{}) rip.Resolver {
	return NewResolver(rip.NewHanlder(source))
}

func (r resolver) Server() rip.Server {
	return NewServer(r)
}
