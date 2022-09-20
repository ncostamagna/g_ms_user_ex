package logger

import (
	"os"

	gkLog "github.com/go-kit/kit/log"
)

type (
	Logger interface {
		Success(msg string)
		Error(msg string)
		Warning(msg string)
		Debug(msg string)
	}

	logger struct {
		gkLog gkLog.Logger
	}
)

func New(gkLog gkLog.Logger) Logger {
	return &logger{
		gkLog: gkLog,
	}
}

func newGoKitLogger(service string) gkLog.Logger {
	logger := gkLog.NewLogfmtLogger(gkLog.NewSyncWriter(os.Stderr))
	return gkLog.With(logger, "ts", gkLog.DefaultTimestampUTC, "service", service)
}

const (
	Success string = "success"
	Error   string = "error"
	Warning string = "warning"
	Debug   string = "debug"
)

func (l *logger) Success(msg string) {
	l.gkLog.Log("status", Success, "message", msg)
}

func (l *logger) Error(msg string) {
	l.gkLog.Log("status", Error, "message", msg)
}

func (l *logger) Warning(msg string) {
	l.gkLog.Log("status", Warning, "message", msg)
}

func (l *logger) Debug(msg string) {
	l.gkLog.Log("status", Debug, "message", msg)
}
