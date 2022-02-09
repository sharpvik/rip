package rfip

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestReader(t *testing.T) {
	input := &Request{
		FuncName: "greet",
		Argument: []byte(`{"hello": "world"}`),
	}

	output, err := ReadRequest(strings.NewReader(input.String()))
	assert.NoError(t, err)
	assert.Equal(t, input.FuncName, output.FuncName)
	assert.Equal(t, input.Argument, output.Argument)
}
