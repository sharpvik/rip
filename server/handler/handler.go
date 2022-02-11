package handler

import (
	"fmt"

	"github.com/sharpvik/rfip"
)

type Handler struct{}

func New() *Handler {
	return new(Handler)
}

func (h *Handler) Greet(name string) string {
	return fmt.Sprintf("Hello, %s, nice to see you here", name)
}

type API struct {
	*rfip.Client
}

func NewAPI(addr string) *API {
	return &API{
		Client: rfip.NewClient(addr),
	}
}

func (i *API) Greet(name string) (resp *rfip.Response) {
	return i.MustInvoke("Greet", name)
}
