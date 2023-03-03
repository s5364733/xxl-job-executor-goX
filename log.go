package xxl

import (
	"fmt"
	"log"
	"os"
)

// LogFunc 应用日志
type LogFunc func(req LogReq, res *LogRes) []byte

// Logger 系统日志
type Logger interface {
	Info(format string, a ...interface{})
	Error(format string, a ...interface{})
	InfoX(logId int64, format string, a ...interface{})
	ErrorX(logId int64, format string, a ...interface{})
}

type Xlogger struct {
}

func (l *Xlogger) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(format, a...))
}

func (l *Xlogger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf(format, a...))
}

func (l *Xlogger) InfoX(logId int64, format string, a ...interface{}) {
	l.logX(logId, format, a)
}

func (l *Xlogger) ErrorX(logId int64, format string, a ...interface{}) {
	l.logX(logId, format, a)
}

func (l *Xlogger) logX(JobId int64, format string, a []interface{}) {
	logStr := fmt.Sprintf(format, a...)
	fmt.Println(logStr)

	filename := fmt.Sprintf("log_%d", JobId)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)

	defer file.Close()

	if err != nil {
		log.Printf("文件创建错误 err: %v", err)
	}
	file.WriteString(logStr + "\n")
}
