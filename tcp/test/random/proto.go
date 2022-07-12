package random

import (
	"math/rand"

	riptcp "github.com/sharpvik/rip/tcp"
)

type Proto struct{}

func NewProto() *Proto {
	return new(Proto)
}

func (p *Proto) RandInt() int {
	return rand.Int()
}

func Server() *riptcp.Server {
	return riptcp.Use(NewProto()).Server()
}
