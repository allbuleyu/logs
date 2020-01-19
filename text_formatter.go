package logs

import (
	"io"
	"runtime"
	"sync"
	"time"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

var baseTimestamp time.Time

func init() {
	baseTimestamp = time.Now()
}

type TextFormatter struct {
	ForceColors bool

	DisableColors bool

	// Override coloring based on CLICOLOR and CLICOLOR_FORCE. - https://bixense.com/clicolors/
	EnvironmentOverrideColors bool

	// Disable timestamp logging. useful when output is redirected to logging
	// system that already adds timestamps.
	DisableTimestamp bool

	FullTimestamp bool

	// default 2006-01-02 15:04:05
	TimestampFormat string

	DisableSorting bool

	SortingFunc func([]string)

	QuoteEmptyFields bool

	isTerminal bool

	FieldMap FieldMap

	// CallerPrettyfier can be set by the user to modify the content
	// of the function and file keys in the data when ReportCaller is
	// activated. If any of the returned value is the empty string the
	// corresponding key will be removed from fields.
	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	terminalInitOnce sync.Once
}

func (t *TextFormatter) init(entry *Entry) {
	if entry.logger != nil {
		t.isTerminal = t.checkIfTerminal(entry.logger.Out)
	}
}

func (t *TextFormatter) checkIfTerminal(w io.Writer) bool {
	return true
}

func (t *TextFormatter) Format(entry *Entry) ([]byte, error) {
	return  nil, nil
}



