package logs

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"time"
)

type Entry struct {
	logger *Logger

	data Fields

	// 创建时间
	cTime time.Time

	caller runtime.Frame

	// 日志缓存
	buf *bytes.Buffer

	ctx context.Context

	err string
}

func NewEntry(l *Logger) *Entry {
	return &Entry{
		logger: l,
		data:   make(Fields, 6),
	}
}

func (entry *Entry) log(level Level, msg string) {

}

func (entry *Entry) Log(level Level, args ...interface{}) {
	if entry.logger.IsLevelEnabled(level) {
		entry.log(level, fmt.Sprint(args...))
	}
}

func (entry *Entry) Print(args ...interface{}) {
	entry.Info(args...)
}

func (entry *Entry) Info(args ...interface{}) {
	entry.Log(InfoLevel, args...)
}

func (entry *Entry) Printf(string, ...interface{}) {
	panic("implement me")
}

func (entry *Entry) Println(...interface{}) {
	panic("implement me")
}

func (entry *Entry) Fatal(...interface{}) {
	panic("implement me")
}

func (entry *Entry) Fatalf(string, ...interface{}) {
	panic("implement me")
}

func (entry *Entry) Fatalln(...interface{}) {
	panic("implement me")
}

func (entry *Entry) Panic(...interface{}) {
	panic("implement me")
}

func (entry *Entry) Panicf(string, ...interface{}) {
	panic("implement me")
}

func (entry *Entry) Panicln(...interface{}) {
	panic("implement me")
}


