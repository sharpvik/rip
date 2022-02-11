package handler

import (
	"fmt"
)

type Handler interface {
	Greet(string) string
}

type handler struct{}

func New() Handler {
	return new(handler)
}

func (h *handler) Greet(name string) string {
	return fmt.Sprintf("Hello, %s, nice to see you here", name)
}
