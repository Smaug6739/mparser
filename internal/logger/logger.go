package logger

import (
	"log"
	"os"
)

const (
	Off     = "\033[0m"
	Green   = "\033[1;32m"
	Yellow  = "\033[1;33m"
	Blue    = "\033[1;34m"
	Red     = "\033[1;31m"
	White   = "\033[1;37m"
	Magenta = "\x1b[35m"
)

type LoggerInterface interface {
	Details(any)
	Info(any)
	Warn(any)
	Error(any)
}

type Logger struct {
	details *log.Logger
	info    *log.Logger
	warn    *log.Logger
	error   *log.Logger
}

func (l *Logger) Details(display any) {
	l.details.Println(displayJson(display))
}

func (l *Logger) Info(str any) {
	l.info.Println(str)
}

func (l *Logger) Warn(str any) {
	l.warn.Println(str)
}

func (l *Logger) Error(str any) {
	l.error.Println(str)
}

func New() LoggerInterface {
	flags := log.Lmsgprefix | log.Ldate | log.Ltime
	return &Logger{
		details: log.New(os.Stdout, Magenta+"[DETAILS]: "+Off, flags),
		info:    log.New(os.Stdout, Blue+"[INFO]: "+Off, flags),
		warn:    log.New(os.Stdout, Yellow+"[WARN]: "+Off, flags),
		error:   log.New(os.Stdout, Red+"[ERROR]: "+Off, flags),
	}
}
