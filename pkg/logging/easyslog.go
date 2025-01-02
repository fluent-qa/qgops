package logging

import (
	"log/slog"
	"os"
)

/**
//https://betterstack.com/community/guides/logging/logging-in-go/
1. Log Level
2. Context-Aware Logging
3. Log Sampling
4. Configuration
5. slog:
	A. Logger
	B. Record
	C. Handler:TextHandler/JSONHandler
*/
/**
set default json log to slog
https://go.dev/blog/slog
*/

var opts = slog.HandlerOptions{
	AddSource: true,
	Level:     slog.LevelDebug,
}

func init() {
	slog.SetDefault(TextLogger())
}
func JsonLogger() *slog.Logger {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opts))
	return logger
}

func TextLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	return logger
}
