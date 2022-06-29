package creational

import (
	"testing"

	"gotest.tools/assert"
)

func TestSingleton(t *testing.T) {
	assert.Equal(t, getInstance().str == Test, true)
}
