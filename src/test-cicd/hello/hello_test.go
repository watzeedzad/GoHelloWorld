package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	actualValue := HelloWorld()
	expectValue := "Hello World"

	assert.Equal(t, actualValue, expectValue)
}
