package logger

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"

	customFormatter.FullTimestamp = true

	logger := logrus.New()
	logger.SetFormatter(customFormatter)

	return logger
}
