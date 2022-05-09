package handler

import (
	"fmt"
	"strings"

	"github.com/sharpvik/rip"
)

type Handler struct{}

func New() *Handler {
	return new(Handler)
}

func (h *Handler) Greet(name string) string {
	return fmt.Sprintf("Hello, %s, nice to see you here", name)
}

func (h *Handler) Users(params UserParams) (users []User) {
	for i := 0; i < params.N; i++ {
		users = append(users, User{
			Name: strings.Repeat("a", params.NameLength),
			Age:  20,
		})
	}
	return
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

func (i *API) Users(params UserParams) (resp *rip.Response) {
	return i.MustInvoke("Users", params)
}
