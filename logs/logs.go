package logs

import (
	"fmt"
	"time"
)

type Level uint8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// Disabled disables the logger.
	Disabled
)

func (l Level) String() string {
	var s string
	switch l {
	case DebugLevel:
		s = "Debug"
	case InfoLevel:
		s = "Info"
	case WarnLevel:
		s = "Warn"
	case ErrorLevel:
		s = "Error"
	case FatalLevel:
		s = "Fatal"
	case PanicLevel:
		s = "Panic"
	case Disabled:
		s = "Disabled"
	default:
		s = "Disabled"
	}

	return s
}

// 90 -97
// https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13
func (l Level) Color() string {
	var s string = "\033[%dm"

	switch l {
	case DebugLevel:
		s = fmt.Sprintf(s, 95)
	case InfoLevel:
		s = fmt.Sprintf(s, 94)
	case WarnLevel:
		s = fmt.Sprintf(s, 93)
	case ErrorLevel:
		s = fmt.Sprintf(s, 91)
	case FatalLevel:
		s = fmt.Sprintf(s, 94)
	case PanicLevel:
		s = fmt.Sprintf(s, 94)
	default:
		s = ""
	}

	return s
}

type Logger struct {
	level Level
}

func (l *Logger) Level(level Level) {
	l.level = level
}

func (l *Logger) Debug(s string) {
	l.level = DebugLevel

	fmt.Printf("%s %s[%s]\033[0m : %s", time.Now().Format("2006-01-02 15:04:05"), l.level.Color(), l.level.String(), s)
}

func (l *Logger) Info(s string) {
	l.level = InfoLevel
	fmt.Printf("%s %s[%s]\033[0m : %s", time.Now().Format("2006-01-02 15:04:05"), l.level.Color(), l.level.String(), s)
}

func (l *Logger) Warn(s string) {
	l.level = WarnLevel
	fmt.Printf("%s %s[%s]\033[0m : %s", time.Now().Format("2006-01-02 15:04:05"), l.level.Color(), l.level.String(), s)
}