package random

import (
	"math/rand"

	"github.com/sharpvik/rip"
	riptcp "github.com/sharpvik/rip/tcp"
)

type Proto struct{}

func NewProto() *Proto {
	return new(Proto)
}

func (p *Proto) RandInt() int {
	return rand.Int()
}

func Server() rip.Server {
	return riptcp.Use(NewProto()).Server()
}
