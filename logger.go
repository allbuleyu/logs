package logs

import (
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)


type Logger struct {
	// 日志输出
	Out io.Writer

	// 是否要开启颜色
	IsColored bool

	// 日志等级
	Level Level

	// 是否打印调用者
	ReportCaller bool

	// 打印格式
	Formatter Formatter

	// 打印字段
	Fields Fields

	// 互斥锁
	mu *sync.Mutex

	// 日志处理池
	entryPool *sync.Pool

	ExitFunc exitFunc
}

type exitFunc func(int)

func NewLog() *Logger {
	return &Logger{
		Out:          os.Stderr,
		IsColored:    false,
		Level:        InfoLevel,
		ReportCaller: false,
		Formatter:    new(TextFormatter),
		Fields:       nil,
		mu:           &sync.Mutex{},
		entryPool:    &sync.Pool{},
	}
}


func (l *Logger) newEntry() *Entry {
	entry, ok := l.entryPool.Get().(*Entry)
	if ok {
		return entry
	}

	return NewEntry(l)
}

func (l *Logger) Logln(args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Logf(args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Log(args ...interface{}) {
	l.log()
}

func (l *Logger) log(args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Print(args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Printf(string, args ...interface{}) {
	fmt.Println(args)
}

func (l *Logger) Println(args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Fatal(args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Fatalf(string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Fatalln(args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Panic(args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Panicf(string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Panicln(args ...interface{}) {
	panic("implement me")
}

func (l *Logger) level() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.Level)))
}

// IsLevelEnabled checks if the log level of the logger is greater than the level param
func (l *Logger) IsLevelEnabled(level Level) bool {
	return l.level() >= level
}