package test

import (
	"rdy24/golang-dependency-injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("test.txt")
	assert.NotNil(t, connection)
	cleanup()
}
