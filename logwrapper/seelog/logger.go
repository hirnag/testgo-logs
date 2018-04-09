package seelog

import (
    "github.com/cihub/seelog"
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
    return &wrapper{}
}

func (wrap wrapper) Debug(msg string) {
    defer seelog.Flush()
    seelog.Debug(msg)
}

func (wrap wrapper) Info(msg string) {
    defer seelog.Flush()
    seelog.Info(msg)
}

func (wrap wrapper) Warn(msg string) {
    defer seelog.Flush()
    seelog.Warn(msg)
}

func (wrap wrapper) Error(msg string) {
    defer seelog.Flush()
    seelog.Error(msg)
}

