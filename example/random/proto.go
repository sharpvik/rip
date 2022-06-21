package random

import (
	"math/rand"
)

type actor struct{}

func Proto() *actor {
	return new(actor)
}

func (p *actor) RandInt() int {
	return rand.Int()
}
