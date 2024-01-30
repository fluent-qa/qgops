package logging

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"testing"
)

func TestSlog(t *testing.T) {
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")
	slog.Error("testing error",
		slog.String("config-test", fmt.Sprintf("%s", "config-value")))
}

func TestJsonHandlerSlog(t *testing.T) {
	logger := JsonLogger()
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")
}

func TestTextHandlerSlog(t *testing.T) {
	logger := TextLogger()
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")
	slog.SetDefault(logger)
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")
}

func TestDefaultSlog(t *testing.T) {
	defaultLogger := log.Default()
	defaultLogger.SetOutput(os.Stdout)
	logger := log.New(os.Stdout, "Test:", log.LstdFlags)
	log.Println("Hello from Go application!")
	logger.Println("Hello from Go application!")
}
