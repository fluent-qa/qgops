package logging

import (
	"testing"
)

func TestNewDefaultZapLogger(t *testing.T) {
	logger := NewDefaultZapLogger()
	logger.Info("test")

}
