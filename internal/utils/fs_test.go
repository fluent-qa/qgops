package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectoryExist(t *testing.T) {
	fs, err := DirectoryExist("./fs.go")
	if err != nil {
		fmt.Println(err)
	}
	fs, _ = DirectoryExist("./tmp")
	assert.NotNil(t, fs)
}
