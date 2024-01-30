package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContainer(t *testing.T) {
	c := NewContainer[any]()
	assert.NotNil(t, c.Injector)
}
