package logging

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"github.com/pquerna/ffjson/ffjson"
)

var (
	log = logrus.New()
)

func init() {
	fmt := &logrus.TextFormatter{
		ForceColors:      true,
		DisableSorting:   true,
		DisableTimestamp: true,
	}

	log.Out = os.Stdout
	log.Formatter = fmt
}

func Output() io.Writer {
	return log.Out
}

func SetOutput(w io.Writer) {
	log.Out = w
}

func SetDebug(enabled bool) {
	if enabled {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
}

func LogScriptError(name string, msg interface{}) {
	log.Errorf("script error [%s]: %v", name, msg)
}

func LogScriptWarn(name string, msg interface{}) {
	log.Warnf("script warn [%s]: %v", name, msg)
}

func Printf(format string, a ...interface{}) {
	fmt.Fprintf(log.Out, format, a...)
}

func Println(a ...interface{}) {
	fmt.Fprintln(log.Out, a...)
}

func Print(a ...interface{}) {
	fmt.Fprint(log.Out, a...)
}

func DDumpUnmarshaled(descr string, in []byte) {
	if log.Level < logrus.DebugLevel {
		return
	}

	var res interface{}
	if err := ffjson.Unmarshal(in, &res); err != nil {
		panic("DumpUnmarshaled: unable to unmarshal input")
	}

	DDump(descr, res)
}

func DDumpJSON(descr string, in interface{}) {
	if log.Level < logrus.DebugLevel {
		return
	}

	DumpJSON(descr, in)
}

func DumpJSON(descr string, in interface{}) {
	out, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		panic("DumpJSON: unable to marshal input")
	}

	Printf("%s ------------------------- dump start ---------------------------------------\n", descr)
	Println(string(out))
	Printf("%s -------------------------  dump end  ---------------------------------------\n\n", descr)
}

func DDump(descr string, in interface{}) {
	if log.Level < logrus.DebugLevel {
		return
	}

	Dump(descr, in)
}

func Dump(descr string, in interface{}) {
	Printf("%s ------------------------- dump start ---------------------------------------\n", descr)
	spew.Fdump(log.Out, in)
	Printf("%s -------------------------  dump end  ---------------------------------------\n\n", descr)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	log.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	log.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	log.Debugln(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	log.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	log.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	log.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	log.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger.
func Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}
