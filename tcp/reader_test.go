package riptcp

import (
	"strings"
	"testing"

	"github.com/sharpvik/rip/proto"
	"github.com/stretchr/testify/assert"
)

func TestRequestReader(t *testing.T) {
	input := &proto.Request{
		Function: "greet",
		Argument: []byte(`{"hello": "world"}`),
	}

	output, err := ReadRequest(strings.NewReader(input.String()))
	assert.NoError(t, err)
	assert.Equal(t, input.Function, output.Function)
	assert.Equal(t, input.Argument, output.Argument)
}
