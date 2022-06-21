package rip

import (
	"fmt"
	"testing"

	"github.com/sharpvik/rip/proto"
	"github.com/stretchr/testify/assert"
)

type api struct{}

type person struct {
	Name string
}

func (i api) Greet(human *person) string {
	return fmt.Sprintf("Hello, %s", human.Name)
}

func TestResolver(t *testing.T) {
	req, err := proto.NewRequest("Greet", person{"Viktor"})
	assert.NoError(t, err)
	assert.Equal(t, proto.ResponseJSON("Hello, Viktor"), Use(new(api)).Handle(req))
}
