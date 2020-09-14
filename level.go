package logger

import (
	"bytes"
	"io"
	"io/ioutil"
	"sync"
)

type LevelFilter interface {
	io.Writer
	Check(line []byte) bool
	SetMinLevel(min LogLevel)
}
type levelFilter struct {
	once      sync.Once
	writer    io.Writer
	badLevels map[LogLevel]struct{}
	levels    []LogLevel
	minLevel  LogLevel
}

func NewLevelFilter(opts ...LevelFilterOption) LevelFilter {
	result := &levelFilter{
		once: sync.Once{},
	}
	for _, opt := range opts {
		opt(result)
	}
	if result.writer == nil {
		result.writer = ioutil.Discard
	}
	if len(result.minLevel) == 0 {
		result.SetMinLevel(InfoLevel)
	}
	if len(result.levels) == 0 {
		result.levels = DefaultLogLevels
	}
	return result
}

// Check will check a given line if it would be included in the level
// filter.
func (l *levelFilter) Check(line []byte) bool {
	l.once.Do(l.init)
	// Check for a log level
	var level LogLevel
	x := bytes.IndexByte(line, '[')
	if x >= 0 {
		y := bytes.IndexByte(line[x:], ']')
		if y >= 0 {
			level = LogLevel(line[x+1 : x+y])
		}
	}

	_, ok := l.badLevels[level]
	return !ok
}

func (l *levelFilter) Write(p []byte) (n int, err error) {
	if !l.Check(p) {
		return len(p), nil
	}
	return l.writer.Write(p)
}

// SetMinLevel is used to update the minimum log level
func (l *levelFilter) SetMinLevel(min LogLevel) {
	l.minLevel = min
	l.init()
}

func (l *levelFilter) init() {
	badLevels := make(map[LogLevel]struct{})
	for _, level := range l.levels {
		if level == l.minLevel {
			break
		}
		badLevels[level] = struct{}{}
	}
	l.badLevels = badLevels
}
