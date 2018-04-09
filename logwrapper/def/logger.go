package def

import (
    "os"
    "log"
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
    writer := os.Stdout
    log.SetOutput(writer)
    return &wrapper{}
}

func (wrap wrapper) Debug(msg string) {
    // nothing debug log
}

func (wrap wrapper) Info(msg string) {
    log.Println(msg)
}

func (wrap wrapper) Warn(msg string) {
    // nothing warn log
}

func (wrap wrapper) Error(msg string) {
    log.Fatalln(msg)
}

