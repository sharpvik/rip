package handler

import (
	"fmt"

	"github.com/sharpvik/rip"
)

type Handler struct{}

func New() *Handler {
	return new(Handler)
}

func (h *Handler) Greet(name string) string {
	return fmt.Sprintf("Hello, %s, nice to see you here", name)
}

type API struct {
	*rip.Client
}

func NewAPI(addr string) *API {
	return &API{
		Client: rip.NewClient(addr),
	}
}

func (i *API) Greet(name string) (resp *rip.Response) {
	return i.MustInvoke("Greet", name)
}
