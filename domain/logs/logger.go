package dominio_logs

import (
	"log"
	"os"
	"sync"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
	loggersOnce sync.Once
)

func InitLoggers() {
	loggersOnce.Do(func() {
		infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		warnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
		errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		fatalLogger = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

func Info(format string, v ...interface{}) {
	InitLoggers()
	infoLogger.Printf(format, v...)
}

func Warn(format string, v ...interface{}) {
	InitLoggers()
	warnLogger.Printf(format, v...)
}

func Error(format string, v ...interface{}) {
	InitLoggers()
	errorLogger.Printf(format, v...)
}

func Fatal(format string, v ...interface{}) {
	InitLoggers()
	fatalLogger.Fatalf(format, v...)
}
