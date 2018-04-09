package logrus

import (
    "github.com/sirupsen/logrus"
    "os"
)

type Logger interface {
    Debug(string)
    Info(string)
    Warn(string)
    Error(string)
}

type wrapper struct {
}

func New() Logger {
    return newlog().(Logger)
}
func newlog() interface{} {
    logrus.SetOutput(os.Stdout)
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetLevel(logrus.DebugLevel)
    return &wrapper{}
}

func (wrap wrapper) Debug(msg string) {
    logrus.Debug(msg)
}

func (wrap wrapper) Info(msg string) {
    logrus.Info(msg)
}

func (wrap wrapper) Warn(msg string) {
    logrus.Warning(msg)
}

func (wrap wrapper) Error(msg string) {
    logrus.Error(msg)
}

