package main

import (
    "log"
    "github.com/cihub/seelog"
    "go.uber.org/zap/zapcore"
    "go.uber.org/zap"
    "github.com/sirupsen/logrus"
    "github.com/golang/glog"

    "time"
    "flag"
    "os"
)

var zaplog *zap.Logger

func main() {

    // -- log output setting for personal

    // -- hello logs
    helloDefault()
    helloGlog()
    helloZap()
    helloLogrus()
    helloSeelog()

    // with parameter
    //url := "http://example.com"
    //num := 10
    //tim := time.Now().UnixNano()
    //helloZapWithParam(url, num, tim)
    //helloLogrusWithParam(url, num, tim)
}

func init() {
    // -- log output setting
    //writer = &bytes.Buffer{}
    //log1, _ := os.Open("log1.log")
    //defer log1.Close()
    //writer = io.MultiWriter(log1)
    writer := os.Stdout

    log.SetOutput(writer)
    flag.Parse()
    logrus.SetOutput(writer)
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetLevel(logrus.DebugLevel)
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
    zaplog, _ = zapconf.Build()
    zaplog, _ = zap.NewDevelopment()
}

func helloDefault() {
    log.Println("Hello from default log")
}

func helloGlog() {
    glog.Error("Hello from Glog")
}

func helloLogrus() {
    logrus.WithFields(logrus.Fields{
    }).Error("Hello from Logrus")
}
func helloLogrusWithParam(url string, num int, tim int64) {
    logrus.WithFields(logrus.Fields{
        "url": url,
        "num":   num,
        "time":   tim,
    }).Error("Hello from Logrus")
}

func helloZap() {
    zaplog, _ = zap.NewProduction()
    zaplog.Error("Hello from Zap",
    )
}
func helloZapWithParam(url string, num int, tim int64) {
    zaplog.Error("Hello from Zap",
        zap.String("url", url),
        zap.Int("num", num),
        zap.Duration("time", time.Duration(tim)),
    )
}

func helloSeelog() {
    defer seelog.Flush()
    seelog.Error("Hello from Seelog")
}

