package rip_test

import (
	"testing"

	"github.com/sharpvik/rip/tcp/test/random"
	"github.com/stretchr/testify/assert"
)

const addr = "localhost:8888"

func TestClientAndServer(t *testing.T) {
	go testServer(t)
	testClient(t)
}

func testServer(t *testing.T) {
	t.Log("server listening at", addr)
	assert.NoError(t, random.Server().ListenAndServeTCP(addr))
}

func testClient(t *testing.T) {
	i, err := random.Client(addr).RandInt()
	assert.NoError(t, err)
	assert.NotEmpty(t, i)
	t.Log("received random number:", i)
}
