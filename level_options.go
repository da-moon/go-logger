package logger

import "io"

// Option - server options setter method
type LevelFilterOption func(*levelFilter)

// string representing log level
type LogLevel string

// default log levels
var (
	TraceLevel LogLevel = "TRACE"
	DebugLevel LogLevel = "DEBUG"
	InfoLevel  LogLevel = "INFO"
	WarnLevel  LogLevel = "WARN"
	ErrorLevel LogLevel = "ERROR"
)
var DefaultLogLevels = []LogLevel{
	TraceLevel,
	DebugLevel,
	InfoLevel,
	WarnLevel,
	ErrorLevel,
}

// WithWriter - sets io.writer
func WithWriter(arg io.Writer) LevelFilterOption {
	return func(s *levelFilter) {
		s.writer = arg
	}
}
func WithMinLevel(arg string) LevelFilterOption {
	return func(s *levelFilter) {
		s.minLevel = LogLevel(arg)
	}
}
func WithLevels(arg []string) LevelFilterOption {
	return func(s *levelFilter) {
		if s.levels == nil {
			s.levels = make([]LogLevel, 0)
		}
		for _, v := range arg {
			s.levels = append(s.levels, LogLevel(v))
		}
	}
}
