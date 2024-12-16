// Package logger defines how logging should occur in a contained enviroment such as a containter
package logger

import (
	"log/slog"
	"os"
)

// Log defines a varable from the slog package
var Log *slog.Logger

// InitLogger intializes the logger with various levels
func InitLogger() {
	// create new logger with json handler
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // default logging level
	})

	// assign to global varible
	Log = slog.New(handler)

	// set the logger as the default logger
	slog.SetDefault(Log)

}

// Info logs an informational message
func Info(msg string, keysAndValues ...interface{}) {
	Log.Info(msg, keysAndValues...)
}

// Warn logs a warning message
func Warn(msg string, keysAndValues ...interface{}) {
	Log.Warn(msg, keysAndValues...)
}

// Error logs an error message
func Error(msg string, err error, keysAndValues ...interface{}) {
	Log.Error(msg, append(keysAndValues, "error", err)...)
}
