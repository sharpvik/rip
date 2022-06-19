package random

import (
	"math/rand"
)

type proto struct{}

func Proto() *proto {
	return new(proto)
}

func (p *proto) RandInt() int {
	return rand.Int()
}
