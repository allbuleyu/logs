package logs

import (
	"io"
	"os"
	"sync"
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
}

func NewLog() *Logger {
	return &Logger{
		Out:          os.Stderr,
		IsColored:    false,
		Level:        InfoLevel,
		ReportCaller: false,
		Formatter:    TextFormatter{},
		Fields:       nil,
		mu:           &sync.Mutex{},
		entryPool:    &sync.Pool{},
	}
}