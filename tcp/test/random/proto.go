package random

import (
	"math/rand"

	riptcp "github.com/sharpvik/rip/tcp"
)

type actor struct{}

func Proto() *actor {
	return new(actor)
}

func (p *actor) RandInt() int {
	return rand.Int()
}

func Server() *riptcp.Server {
	return riptcp.Use(Proto()).Server()
}
