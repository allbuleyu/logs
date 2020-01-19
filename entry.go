package logs

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	bufferPool *sync.Pool
)

type Entry struct {
	logger *Logger

	level Level

	message string

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
	var buffer *bytes.Buffer

	if entry.cTime.IsZero() {
		entry.cTime = time.Now()
	}

	entry.level = level
	entry.message = msg

	buffer = bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer bufferPool.Put(buffer)
	entry.buf = buffer

	entry.write()

	entry.buf = nil

	if level <= PanicLevel {
		panic(&entry)
	}
}

func (entry *Entry) write() {
	entry.logger.mu.Lock()
	defer entry.logger.mu.Unlock()

	serialized, err := entry.logger.Formatter.Format(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
		return
	}

	_, err = entry.logger.Out.Write(serialized)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to wirte to log, %v\n", err)
	}
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


