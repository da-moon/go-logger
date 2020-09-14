package logger

import (
	"io"
	"log"
	"strings"
)

// WrappedLogger is a logger that adds default log level
// prefixes and exposes methods for writing to
// Those levels
type WrappedLogger struct {
	logger *log.Logger
}

// NewWrappedLogger returns a new Wrapped Logger
func NewWrappedLogger(l *log.Logger) *WrappedLogger {
	if l == nil {
		panic("logger was nil")
	}
	result := &WrappedLogger{
		logger: l,
	}
	return result
}

// Trace ...
func (l *WrappedLogger) Trace(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, "[TRACE]")
	format = strings.TrimSpace(format)
	format = "[TRACE] " + format
	l.logger.Printf(format, v...)
}

// Debug ...
func (l *WrappedLogger) Debug(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, "[DEBUG]")
	format = strings.TrimSpace(format)
	format = "[DEBUG] " + format
	l.logger.Printf(format, v...)
}

// Info ...
func (l *WrappedLogger) Info(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, "[INFO]")
	format = strings.TrimSpace(format)
	format = "[INFO] " + format
	l.logger.Printf(format, v...)
}

// Warn ...
func (l *WrappedLogger) Warn(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, "[WARN]")
	format = strings.TrimSpace(format)
	format = "[WARN] " + format
	l.logger.Printf(format, v...)
}

// Error ...
func (l *WrappedLogger) Error(format string, v ...interface{}) {
	format = strings.TrimPrefix(format, "[ERROR]")
	format = strings.TrimSpace(format)
	format = "[ERROR] " + format
	l.logger.Printf(format, v...)
}

// Output ...
func (l *WrappedLogger) Output(calldepth int, s string) error {
	return l.logger.Output(calldepth, s)
}

// Printf ...
func (l *WrappedLogger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

// Print ...
func (l *WrappedLogger) Print(v ...interface{}) {
	l.logger.Print(v...)
}

// Println ...
func (l *WrappedLogger) Println(v ...interface{}) {
	l.logger.Println(v...)
}

// Fatal ...
func (l *WrappedLogger) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}

// Fatalf ...
func (l *WrappedLogger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

// Fatalln ...
func (l *WrappedLogger) Fatalln(v ...interface{}) {
	l.logger.Fatalln(v...)
}

// Panic ...
func (l *WrappedLogger) Panic(v ...interface{}) {
	l.logger.Panic(v...)
}

// Panicf ...
func (l *WrappedLogger) Panicf(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}

// Panicln ...
func (l *WrappedLogger) Panicln(v ...interface{}) {
	l.logger.Panicln(v...)
}

// Flags ...
func (l *WrappedLogger) Flags() int {
	return l.logger.Flags()
}

// SetFlags ...
func (l *WrappedLogger) SetFlags(flag int) {
	l.logger.SetFlags(flag)
}

// Prefix ...
func (l *WrappedLogger) Prefix() string {
	return l.logger.Prefix()
}

// SetPrefix ...
func (l *WrappedLogger) SetPrefix(prefix string) {
	l.logger.SetPrefix(prefix)
}

// Writer ...
func (l *WrappedLogger) Writer() io.Writer {
	return l.logger.Writer()
}
