package logging

import (
	"fmt"

	"github.com/rollbar/rollbar-go"
	cons "github.com/carcinodehyde/OCR-Go/constant"

	"os"

	opLogging "github.com/op/go-logging"
)

var log = opLogging.MustGetLogger(cons.LOG_MODULE)

const INTERNAL = "internal"

func MustGetLogger(name string) *AppLogger {
	host, err := os.Hostname()
	if err != nil {
		log.Error("", err.Error())
		host = "unknown"
	}
	appl := &AppLogger{opLogging.MustGetLogger(name), host}
	appl.ExtraCalldepth = 1
	return appl
}

type AppLogger struct {
	*opLogging.Logger
	Hostname string
}

func (appl *AppLogger) Debug(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + appl.Hostname + "] [" + userid + "]"}, args...)
	appl.Logger.Debug(args...)
}

func (appl *AppLogger) Debugf(userid string, string_format string, args ...interface{}) {
	appl.Logger.Debugf("["+appl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (appl *AppLogger) Info(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + appl.Hostname + "] [" + userid + "]"}, args...)
	appl.Logger.Info(args)
}

func (appl *AppLogger) Infof(userid string, string_format string, args ...interface{}) {
	appl.Logger.Infof("["+appl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (appl *AppLogger) Error(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + appl.Hostname + "] [" + userid + "]"}, args...)
	appl.Logger.Error(args)
}

func (appl *AppLogger) Errorf(userid string, string_format string, args ...interface{}) {
	appl.Logger.Errorf("["+appl.Hostname+"] ["+userid+"] "+string_format, args...)
	rollbar.Error("[" + appl.Hostname + "] [" + userid + "] " + fmt.Sprintf(string_format, args...))
}

func (appl *AppLogger) Critical(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + appl.Hostname + "] [" + userid + "]"}, args...)
	appl.Logger.Critical(args)
}

func (appl *AppLogger) Criticalf(userid string, string_format string, args ...interface{}) {
	appl.Logger.Criticalf("["+appl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (appl *AppLogger) Fatal(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + appl.Hostname + "] [" + userid + "]"}, args...)
	appl.Logger.Fatal(args)
}

func (appl *AppLogger) Fatalf(userid string, string_format string, args ...interface{}) {
	appl.Logger.Fatalf("["+appl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (appl *AppLogger) Panic(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + appl.Hostname + "] [" + userid + "]"}, args...)
	appl.Logger.Panic(args)
}

func (appl *AppLogger) Panicf(userid string, string_format string, args ...interface{}) {
	appl.Logger.Panicf("["+appl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (appl *AppLogger) Warning(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + appl.Hostname + "] [" + userid + "]"}, args...)
	appl.Logger.Warning(args)
}

func (appl *AppLogger) Warningf(userid string, string_format string, args ...interface{}) {
	appl.Logger.Warningf("["+appl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (appl *AppLogger) Notice(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + appl.Hostname + "] [" + userid + "]"}, args...)
	appl.Logger.Notice(args)
}

func (appl *AppLogger) Noticef(userid string, string_format string, args ...interface{}) {
	appl.Logger.Noticef("["+appl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func init() {

	format := opLogging.MustStringFormatter(
		`%{color} %{time:2006-01-02T15:04:05.999Z07:00} %{shortfile:20.20s} %{shortfunc:20.20s} ▶ %{level:.4s} %{message}%{color:reset}`,
	)
	backend := opLogging.NewLogBackend(os.Stderr, "", 0)
	formatter := opLogging.NewBackendFormatter(backend, format)
	opLogging.SetBackend(formatter)
	log.Info(INTERNAL, "logging initialized")
}
