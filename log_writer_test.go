package logger_test

import (
	"testing"

	logger "github.com/da-moon/go-logger"
	assert "github.com/stretchr/testify/assert"
)

type MockLogHandler struct {
	logs []string
}

func (m *MockLogHandler) HandleLog(l string) {
	m.logs = append(m.logs, l)
}

func TestLogWriter(t *testing.T) {
	h := &MockLogHandler{}
	w := logger.NewLogWriter(4)

	w.Write([]byte("one"))
	w.Write([]byte("two"))
	w.Write([]byte("three"))
	w.Write([]byte("four"))
	w.Write([]byte("five"))

	w.RegisterHandler(h)

	w.Write([]byte("six"))
	w.Write([]byte("seven"))

	w.DeregisterHandler(h)

	w.Write([]byte("eight"))
	w.Write([]byte("nine"))

	out := []string{
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
	}
	for idx := range out {
		assert.Equal(t, out[idx], h.logs[idx])
	}
}
