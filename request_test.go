package rfip

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestReader(t *testing.T) {
	input := &Request{
		URL: strings.Split("url/path/to/specify/function/to/call", "/"),
		Arg: []byte(`{"hello": "world"}`),
	}

	output, err := ReadRequest(strings.NewReader(input.String()))
	assert.NoError(t, err)
	assert.Equal(t, input.URL, output.URL)
	assert.Equal(t, input.Arg, output.Arg)
}
