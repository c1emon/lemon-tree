package log

import (
	"io"
	"sync"

	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

var el *echoLogrus
var onceEcho = sync.Once{}

// echoLogrus extend logrus.Logger
type echoLogrus struct {
	*logrus.Logger
}

func GetEchoLogrusLogger() *echoLogrus {
	onceEcho.Do(func() {
		el = &echoLogrus{GetLogger()}
	})
	return el
}

// Echo2LogrusLogLevel to logrus.Level
func Echo2LogrusLogLevel(level log.Lvl) logrus.Level {
	switch level {
	case log.DEBUG:
		return logrus.DebugLevel
	case log.INFO, log.OFF:
		return logrus.InfoLevel
	case log.WARN:
		return logrus.WarnLevel
	case log.ERROR:
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}

}

// Logrus2EchoLogLevel to Echo.log.lvl
func Logrus2EchoLogLevel(level logrus.Level) log.Lvl {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return log.ERROR
	default:
		return log.OFF
	}
}

// Output return log io.Writer
func (l *echoLogrus) Output() io.Writer {
	return l.Out
}

// SetOutput log io.Writer
func (l *echoLogrus) SetOutput(w io.Writer) {
	l.Out = w
}

// Level return log level
func (l *echoLogrus) Level() log.Lvl {
	return Logrus2EchoLogLevel(l.Logger.Level)
}

// SetLevel log level
func (l *echoLogrus) SetLevel(v log.Lvl) {
	l.Logger.Level = Echo2LogrusLogLevel(v)
}

// SetHeader log header
// Managed by Logrus itself
// This function do nothing
func (l *echoLogrus) SetHeader(h string) {
	// do nothing
}

// Prefix return log prefix
// This function do nothing
func (l *echoLogrus) Prefix() string {
	return ""
}

// SetPrefix log prefix
// This function do nothing
func (l *echoLogrus) SetPrefix(p string) {
	// do nothing
}

// Print output message of print level
func (l *echoLogrus) Print(i ...interface{}) {
	l.Logger.Print(i...)
}

// Printf output format message of print level
func (l *echoLogrus) Printf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

// Printj output json of print level
func (l *echoLogrus) Printj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Println(string(b))
}

// Debug output message of debug level
func (l *echoLogrus) Debug(i ...interface{}) {
	l.Logger.Debug(i...)
}

// Debugf output format message of debug level
func (l *echoLogrus) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

// Debugj output message of debug level
func (l *echoLogrus) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Debugln(string(b))
}

// Info output message of info level
func (l *echoLogrus) Info(i ...interface{}) {
	l.Logger.Info(i...)
}

// Infof output format message of info level
func (l *echoLogrus) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

// Infoj output json of info level
func (l *echoLogrus) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Infoln(string(b))
}

// Warn output message of warn level
func (l *echoLogrus) Warn(i ...interface{}) {
	l.Logger.Warn(i...)
}

// Warnf output format message of warn level
func (l *echoLogrus) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

// Warnj output json of warn level
func (l *echoLogrus) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Warnln(string(b))
}

// Error output message of error level
func (l *echoLogrus) Error(i ...interface{}) {
	l.Logger.Error(i...)
}

// Errorf output format message of error level
func (l *echoLogrus) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

// Errorj output json of error level
func (l *echoLogrus) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Errorln(string(b))
}

// Fatal output message of fatal level
func (l *echoLogrus) Fatal(i ...interface{}) {
	l.Logger.Fatal(i...)
}

// Fatalf output format message of fatal level
func (l *echoLogrus) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

// Fatalj output json of fatal level
func (l *echoLogrus) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Fatalln(string(b))
}

// Panic output message of panic level
func (l *echoLogrus) Panic(i ...interface{}) {
	l.Logger.Panic(i...)
}

// Panicf output format message of panic level
func (l *echoLogrus) Panicf(format string, args ...interface{}) {
	l.Logger.Panicf(format, args...)
}

// Panicj output json of panic level
func (l *echoLogrus) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Panicln(string(b))
}
