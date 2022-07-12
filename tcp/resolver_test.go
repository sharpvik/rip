package riptcp

import (
	"fmt"
	"testing"

	"github.com/sharpvik/rip"
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
	req, e := rip.NewRequest("Greet", person{"Viktor"})
	assert.NoError(t, e)
	assert.Equal(t, rip.ResponseJSON("Hello, Viktor"), Use(new(api)).Handle(req))
}
