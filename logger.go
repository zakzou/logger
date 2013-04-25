package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	CRITICAL
)

type Logger struct {
	log    *log.Logger
	level  int
	caller int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewLogger(filename string, level int, caller int) *Logger {
	var output *os.File
	var err error
	if filename == "" {
		output = os.Stderr
	} else {
		output, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		checkErr(err)
	}
	logger := new(Logger)
	logger.level = level
	logger.log = log.New(output, "", log.Ldate|log.Ltime)
	logger.caller = caller
	return logger
}

func (logger *Logger) SetLevel(level int) {
	logger.level = level
}

func (logger *Logger) GetLevel() int {
	return logger.level
}

func (logger *Logger) print(level int, prefix string, v ...interface{}) {
	if logger.level <= level {
		_, filepath, line, _ := runtime.Caller(logger.caller)
		filename := path.Base(filepath)
		fileinfo := fmt.Sprintf("file: %s, line: %d ", filename, line)
		errorinfo := fmt.Sprint(v...)
		logger.log.Println(prefix, fileinfo, errorinfo)
	}
}

func (logger *Logger) Debug(v ...interface{}) {
	logger.print(DEBUG, "[DEBUG]", v...)
}

func (logger *Logger) Info(v ...interface{}) {
	logger.print(INFO, "[INFO]", v...)
}

func (logger *Logger) Warning(v ...interface{}) {
	logger.print(WARNING, "[WARNING]", v...)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.print(ERROR, "[ERROR]", v...)
}
