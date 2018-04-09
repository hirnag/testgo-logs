package main

import (
    "github.com/hirnag/testgo-logs/logwrapper/def"
    "github.com/hirnag/testgo-logs/logwrapper/zap"
    "github.com/hirnag/testgo-logs/logwrapper/seelog"
    "github.com/hirnag/testgo-logs/logwrapper/glog"
    "github.com/hirnag/testgo-logs/logwrapper/logrus"

    "fmt"
)

var deflogger def.Logger
var zaplogger zap.Logger
var seelogger seelog.Logger
var glogger glog.Logger
var lrlogger logrus.Logger
var bigv string = "Hello from "

func main() {

    // -- log output setting for personal
    Setup()

    // -- hello logs
    helloDefault()
    helloGlog()
    helloLogrus()
    helloZap()
    helloSeelog()

}

func Setup() {

    // logger initialize with to wrapped
    deflogger = def.New()
    glogger = glog.New()
    lrlogger = logrus.New()
    zaplogger = zap.New()
    seelogger = seelog.New()

    for i := 0; i < 50; i++ {
        bigv += "qazwsxedcrfvtgbyhnuj"
    }
    fmt.Println(len(bigv))
}

// default logging
func helloDefault() {
    deflogger.Info("Hello from default log")
}
func bigvalueDefault() {
    deflogger.Info(bigv)
}

// glog
func helloGlog() {
    glogger.Error("Hello from Glog")
}
func bigvalueGlog() {
    glogger.Error(bigv)
}

// logrus
func helloLogrus() {
    lrlogger.Info("Hello from Logrus")
}
func bigvalueLogrus() {
    lrlogger.Info(bigv)
}

// zap
func helloZap() {
    zaplogger.Info("Hello from Zap")
}
func bigvalueZap() {
    zaplogger.Info(bigv)
}

// seelog
func helloSeelog() {
    seelogger.Error("Hello from Seelog")
}
func bigvalueSeelog() {
    seelogger.Error(bigv)
}

//func helloLogrusWithParam(url string, num int, tim int64) {
//    logrus.WithFields(logrus.Fields{
//        "url": url,
//        "num":   num,
//        "time":   tim,
//    }).Error("Hello from Logrus")
//}
//func helloZapWithParam(url string, num int, tim int64) {
//    zaplog.Error("Hello from Zap",
//        zap.String("url", url),
//        zap.Int("num", num),
//        zap.Duration("time", time.Duration(tim)),
//    )
//}
