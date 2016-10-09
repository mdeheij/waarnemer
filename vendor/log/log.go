package log

import (
	"os"

	"github.com/op/go-logging"
)

var stdoutFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05}: %{message}`,
)

// var logfileFormat = logging.MustStringFormatter(
// 	`%{color}%{time:15:04:05} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
// )

//Log is an exported logging variable
var Log = logging.MustGetLogger("example")

func init() {
	// logfile := logging.NewLogBackend("", "", 0)
	stdout := logging.NewLogBackend(os.Stdout, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	// backend1Formatter := logging.NewBackendFormatter(logfile, format)
	backend2Formatter := logging.NewBackendFormatter(stdout, stdoutFormat)

	// Only errors and more severe messages should be sent to backend1
	//backend1Leveled := logging.AddModuleLevel(backend1)
	//backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	// logging.SetBackend(backend1Formatter, backend2Formatter)
	logging.SetBackend(backend2Formatter)

	// Log.Info("Initialized logging.")
}

//Debugf  See Log.Debugf
func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

//Debug  See Log.Debug
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

//Info  See Log.Info
func Info(args ...interface{}) {
	Log.Info(args...)
}

//Println  See Log.Info
func Println(args ...interface{}) {
	Log.Info(args...)
}

//Notice  See Log.Notice
func Notice(args ...interface{}) {
	Log.Notice(args...)
}

//Warning  See Log.Warning
func Warning(args ...interface{}) {
	Log.Warning(args...)
}

//Error  See Log.Error
func Error(args ...interface{}) {
	Log.Error(args...)
}

//Critical  See Log.Critical
func Critical(args ...interface{}) {
	Log.Critical(args...)
}

//Panic  See Log.Panic
func Panic(args ...interface{}) {
	Log.Panic(args...)
}

//Fatal  See Log.Fatal
func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}
