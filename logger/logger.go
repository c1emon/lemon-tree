package logger

import (
	"fmt"
	"github.com/c1emon/lemontree/util"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var cstZone = time.FixedZone("GMT", 8*3600)

// CostumeLogFormatter Custom log format definition
type costumeLogFormatter struct{}

// Format log format
func (s *costumeLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var colorFormater func(a ...interface{}) string
	switch entry.Level {
	case logrus.DebugLevel:
		colorFormater = color.New(color.FgHiYellow).SprintFunc()
	case logrus.InfoLevel:
		colorFormater = color.New(color.FgGreen).SprintFunc()
	case logrus.WarnLevel:
		colorFormater = color.New(color.FgYellow).SprintFunc()
	default:
		colorFormater = color.New(color.FgRed).SprintFunc()
	}

	timestamp := time.Now().In(cstZone).Format("2006-01-02 15:04:05.999")
	msg := fmt.Sprintf("%s [%s] -- %s\n",
		timestamp,
		colorFormater(strings.ToUpper(entry.Level.String())),
		entry.Message)
	if entry.Data != nil && len(entry.Data) > 0 {
		msg = fmt.Sprintf("%s\n%s\n", msg, util.PrettyMarshal(entry.Data))
	}

	return []byte(msg), nil
}

func Init(level string) {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetFormatter(new(costumeLogFormatter))
	logrus.SetLevel(lv)
	logrus.Info(fmt.Sprintf("log level: %s", logrus.GetLevel().String()))
}

// TODO: impl below with logrus
//Logger interface {
//Output() io.Writer
//SetOutput(w io.Writer)
//Prefix() string
//SetPrefix(p string)
//Level() log.Lvl
//SetLevel(v log.Lvl)
//SetHeader(h string)
//Print(i ...interface{})
//Printf(format string, args ...interface{})
//Printj(j log.JSON)
//Debug(i ...interface{})
//Debugf(format string, args ...interface{})
//Debugj(j log.JSON)
//Info(i ...interface{})
//Infof(format string, args ...interface{})
//Infoj(j log.JSON)
//Warn(i ...interface{})
//Warnf(format string, args ...interface{})
//Warnj(j log.JSON)
//Error(i ...interface{})
//Errorf(format string, args ...interface{})
//Errorj(j log.JSON)
//Fatal(i ...interface{})
//Fatalj(j log.JSON)
//Fatalf(format string, args ...interface{})
//Panic(i ...interface{})
//Panicj(j log.JSON)
//Panicf(format string, args ...interface{})
//}

type LogBridge struct {
}

func (l *LogBridge) Write(p []byte) (n int, err error) {
	logrus.Info(string(p))
	return len(p), nil
}
