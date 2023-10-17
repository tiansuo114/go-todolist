package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type logger struct {
	logger *logrus.Logger
	entry  *logrus.Entry
}

func newLogger() logger {
	Logger := logger{logger: logrus.New()}
	Logger.logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 将日志同时输出到文件和控制台
	logFilePath := "logs/"
	now := time.Now()
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err == nil {
		Logger.logger.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		log.Fatalf("Failed to open log file %s: %v", fileName, err)
	}

	Logger.getCallerInfo()
	return Logger
}

var l *logger

func init() {

	_, err := init_OutputFile()
	if err != nil {
		log.Fatalf("failed to set output file: %v", err)
	}

	logger_stu := newLogger()

	l = &logger_stu
}

func init_OutputFile() (*os.File, error) {
	// 设置日志文件路径
	now := time.Now()
	logFilePath := "logs/"
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logFilePath, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to create log directory: %v", err)
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	// 创建日志文件
	_, err = os.Stat(fileName)
	file := &os.File{}
	if os.IsNotExist(err) {
		file, err = os.Create(fileName)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	} else {
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	return file, nil
}

func (l *logger) getCallerInfo() {
	// 获取调用栈,找到调用WriteLog的函数信息
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	entry := l.logger.WithFields(logrus.Fields{
		"funcName": path.Base(funcName),
		"file":     file,
		"line":     line,
	})
	l.entry = entry
}

func WriteLog(err string) {
	level := l.getLevel(err)
	l.getCallerInfo()
	entry := l.entry.WithFields(logrus.Fields{
		"funcName": l.entry.Data["funcName"],
		"file":     l.entry.Data["file"],
		"line":     l.entry.Data["line"],
	})
	entry.Log(level, err) // 直接使用WithFields的结果输出日志

}

func (l *logger) getLevel(err string) logrus.Level {
	if strings.Contains(err, "[Emergency]") {
		return logrus.PanicLevel
	} else if strings.Contains(err, "[Alert]") {
		return logrus.FatalLevel
	} else if strings.Contains(err, "[Critical]") {
		return logrus.ErrorLevel
	} else if strings.Contains(err, "[Error]") {
		return logrus.ErrorLevel
	} else if strings.Contains(err, "[Warning]") {
		return logrus.WarnLevel
	} else if strings.Contains(err, "[Notice]") {
		return logrus.InfoLevel
	} else {
		return logrus.DebugLevel
	}
}

func (l *logger) crashHandle() {
	if err := recover(); err != nil {
		l.logger.WithFields(logrus.Fields{
			"funcName": "crashHandle",
		}).Panic(err)
	}
}
