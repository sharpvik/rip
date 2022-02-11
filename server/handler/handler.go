package handler

import (
	"fmt"

	"github.com/sharpvik/rfip"
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

type api struct {
	*rfip.Client
}

func NewAPI(addr string) Handler {
	return &api{
		Client: rfip.NewClient(addr),
	}
}

func (i *api) Greet(name string) (s string) {
	i.MustInvoke("Greet", name).MustUnmarshal(&s)
	return
}
