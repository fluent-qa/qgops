package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileExtension(t *testing.T) {
	result := GetFileExtension("a.json")
	assert.Equal(t, "json", result)
}
