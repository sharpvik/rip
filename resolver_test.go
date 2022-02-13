package rip

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExampleAPI struct{}
type ExampleArg struct {
	Name string `json:"name"`
}

func (api ExampleAPI) Greet(arg *ExampleArg) string {
	return fmt.Sprintf("Hello, %s", arg.Name)
}

func TestResolver(t *testing.T) {
	resr := NewResolver(ExampleAPI{})
	req, err := NewRequest("Greet", ExampleArg{"Viktor"})
	assert.NoError(t, err)
	resp := resr.Resolve(req)
	expect := ResponseJSON("Hello, Viktor")
	assert.Equal(t, expect, resp)
}
