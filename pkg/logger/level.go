package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (l *Logger) Debug(message string, fields Fields) {
	l.logger.WithFields(fields).Debug(message)
}

func (l *Logger) Error(message string, err error, fields Fields) {
	if err == nil {
		l.logger.WithFields(fields).Error(message)
	} else {
		l.logger.WithFields(fields).Error(fmt.Sprintf("%s: %s", message, err.Error()))
	}
}

func (l *Logger) Info(message string, fields Fields) {
	l.logger.WithFields(fields).Info(message)
}

func (l *Logger) Text(message string) {
	l.logger.Info(message)
}

func (l *Logger) Print(message string, data interface{}) {
	l.logger.WithFields(logrus.Fields{
		"data": fmt.Sprintf("%+v", data),
	}).Info(message)
}
