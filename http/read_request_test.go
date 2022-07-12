package riphttp

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestReader(t *testing.T) {
	argument := `{"name": "Viktor"}`
	r, err := http.NewRequest(http.MethodGet, "/Greet", strings.NewReader(argument))
	assert.NoError(t, err)
	req, e := ReadRequest(r)
	assert.NoError(t, e)
	assert.Equal(t, req.Function, "Greet")
	assert.Equal(t, req.Argument, []byte(argument))
}
