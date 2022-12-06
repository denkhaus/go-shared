package zaplog

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

type LoggerSeverity int

const (
	LoggerSeverityInfo LoggerSeverity = iota
	LoggerSeverityError
)

var (
	zapLogger *zap.Logger
)

func init() {
	zl, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	zapLogger = zl
}

type PrintfLogger struct {
	severity LoggerSeverity
}

func (p *PrintfLogger) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)

	switch p.severity {
	case LoggerSeverityInfo:
		zapLogger.Info(msg)
	case LoggerSeverityError:
		zapLogger.Error(msg)
	}
}

func Logger() *zap.Logger {
	return zapLogger
}

func InfoLogger() *PrintfLogger {
	return &PrintfLogger{
		severity: LoggerSeverityInfo,
	}
}

func ErrorLogger() *PrintfLogger {
	return &PrintfLogger{
		severity: LoggerSeverityError,
	}
}
