package logger

import (
	"fmt"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Init(level string) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logger.Warn(fmt.Sprintf("Invalid log level '%s'; using default level 'info'", level))
		logLevel = logrus.InfoLevel // Default level
	}
	logger.SetLevel(logLevel)
	logger.SetFormatter(&nested.Formatter{
		HideKeys: true,
	})
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func Debug(message string) {
	logger.Debug(message)
}

func Info(message string) {
	logger.Info(message)
}

func Warn(message string) {
	logger.Warn(message)
}

func Error(err error) {
	logger.Error(err.Error())
}

func Panic(err error) {
	logger.Panic(err.Error())
}
