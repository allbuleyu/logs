package logs

import "fmt"

type Level uint32

func (l Level) String() string {
	if lvl,err := l.MarshalText(); err == nil {
		return string(lvl)
	}

	return "unknown"
}

func (l *Level) UnmarshalText(text []byte) error {
	lvl, err := ParseLevel(string(text))
	if err != nil {
		return err
	}

	*l = lvl

	return nil
}

func (l Level) MarshalText() (text []byte, err error) {
	switch l {
	case PanicLevel:
		return []byte("panic"), nil
	case FatalLevel:
		return []byte("fatal"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case WarnLevel:
		return []byte("warn"), nil
	case InfoLevel:
		return []byte("info"), nil
	case DebugLevel:
		return []byte("debug"), nil
	case TraceLevel:
		return []byte("trace"), nil
	}

	return nil, fmt.Errorf("not a valid logs level %d", l)
}

func ParseLevel(level string) (Level, error) {
	switch level {
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "error":
		return ErrorLevel, nil
	case "warn", "waring":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	case "trace":
		return TraceLevel, nil
	}

	return 0, fmt.Errorf("not a valid logs Level: %q", level)
}

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

type Fields map[string]string


type StdLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}