package riptcp

import (
	"strings"
	"testing"

	"github.com/sharpvik/rip"
	"github.com/stretchr/testify/assert"
)

func TestRequestReader(t *testing.T) {
	input := &rip.Request{
		Function: "greet",
		Argument: []byte(`{"hello": "world"}`),
	}

	output, e := ReadRequest(strings.NewReader(input.String()))
	assert.NoError(t, e)
	assert.Equal(t, input.Function, output.Function)
	assert.Equal(t, input.Argument, output.Argument)
}
