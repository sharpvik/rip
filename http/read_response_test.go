package riphttp

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadResponse(t *testing.T) {
	body := []byte("hello world")
	resp := ReadResponse(&http.Response{
		Body: io.NopCloser(bytes.NewReader(body)),
	})
	assert.True(t, resp.OK())
	assert.Equal(t, len(body), resp.Len)
	assert.Equal(t, body, resp.Body)
}
