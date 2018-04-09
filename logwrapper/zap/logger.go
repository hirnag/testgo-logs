package zap

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

type Logger interface {
    Debug(string)
    Info(string)
    Warn(string)
    Error(string)
}

type wrapper struct {
    logger *zap.Logger
}

func New() Logger {
    return newlog().(Logger)
}
func newlog() interface{} {
    zapconf := zap.Config{
        Level: zap.NewAtomicLevelAt(zapcore.DebugLevel),
        Encoding: "json",
        EncoderConfig: zapcore.EncoderConfig{
            TimeKey:        "Time",
            LevelKey:       "Level",
            NameKey:        "Name",
            CallerKey:      "Caller",
            MessageKey:     "Msg",
            StacktraceKey:  "St",
            EncodeLevel:    zapcore.CapitalLevelEncoder,
            EncodeTime:     zapcore.ISO8601TimeEncoder,
            EncodeDuration: zapcore.StringDurationEncoder,
            EncodeCaller:   zapcore.ShortCallerEncoder,
        },
        OutputPaths: []string{"stdout"},
        ErrorOutputPaths: []string{"stderr"},
    }
    lo, err := zapconf.Build()
    if err != nil { panic(err) }

    return &wrapper{ lo }
}

func (wrap wrapper) Debug(msg string) {
    wrap.logger.Debug(msg)
}

func (wrap wrapper) Info(msg string) {
    wrap.logger.Info(msg)
}

func (wrap wrapper) Warn(msg string) {
    wrap.logger.Warn(msg)
}

func (wrap wrapper) Error(msg string) {
    wrap.logger.Error(msg)
}

