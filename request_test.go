package rip

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestReader(t *testing.T) {
	input := &Request{
		Function: "greet",
		Argument: []byte(`{"hello": "world"}`),
	}

	output, err := ReadRequest(strings.NewReader(input.String()))
	assert.NoError(t, err)
	assert.Equal(t, input.Function, output.Function)
	assert.Equal(t, input.Argument, output.Argument)
}
