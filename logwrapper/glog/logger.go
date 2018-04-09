package glog

import (
    "github.com/golang/glog"
    "flag"
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
    // nothing debug log
}

func (wrap wrapper) Info(msg string) {
    flag.Parse()
    glog.Info(msg)
}

func (wrap wrapper) Warn(msg string) {
    flag.Parse()
    glog.Warning(msg)
}

func (wrap wrapper) Error(msg string) {
    flag.Parse()
    glog.Error(msg)
}

