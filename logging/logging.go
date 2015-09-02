package logging

import (
    "fmt"
    "log"
    "io"
//    "io/ioutil"
//    "os"
)

var (
    debug   *log.Logger
    info    *log.Logger
    warn    *log.Logger
    error   *log.Logger
    fatal   *log.Logger
)

func Init(
    debugHandle io.Writer,
    infoHandle io.Writer,
    warnHandle io.Writer,
    errorHandle io.Writer,
    fatalHandle io.Writer) {

    debug = log.New(debugHandle,
        "DEBUG: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    info = log.New(infoHandle,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    warn = log.New(warnHandle,
        "WARN: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    error = log.New(errorHandle,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    fatal = log.New(fatalHandle,
        "FATAL: ",
        log.Ldate|log.Ltime|log.Lshortfile)
}

func Debug(log_string string, a ...interface{}) {
    debug.Printf(fmt.Sprintf("%v\n", log_string), a...)
}

func Info(log_string string, a ...interface{}) {
    info.Printf(fmt.Sprintf("%v\n", log_string), a...)
}

func Warn(log_string string, a ...interface{}) {
    warn.Printf(fmt.Sprintf("%v\n", log_string), a...)
}

func Error(log_string string, a ...interface{}) {
    error.Printf(fmt.Sprintf("%v\n", log_string), a...)
}

func Fatal(log_string string, a ...interface{}) {
    fatal.Printf(fmt.Sprintf("%v\n", log_string), a...)
}

