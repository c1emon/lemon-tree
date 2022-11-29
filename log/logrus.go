package log

import (
	"fmt"
	"github.com/c1emon/lemontree/util"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func GetLogger() *logrus.Logger {
	return logger
}

var cstZone = time.FixedZone("GMT", 8*3600)

// CostumeLogFormatter Custom log format definition
type costumeLogFormatter struct{}

// Format log format
func (s *costumeLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var colorFormatter func(a ...interface{}) string
	switch entry.Level {
	case logrus.DebugLevel:
		colorFormatter = color.New(color.FgHiYellow).SprintFunc()
	case logrus.InfoLevel:
		colorFormatter = color.New(color.FgGreen).SprintFunc()
	case logrus.WarnLevel:
		colorFormatter = color.New(color.FgYellow).SprintFunc()
	default:
		colorFormatter = color.New(color.FgRed).SprintFunc()
	}

	timestamp := time.Now().In(cstZone).Format("2006-01-02 15:04:05.999")
	msg := fmt.Sprintf("%s [%s] -- %s\n",
		timestamp,
		colorFormatter(strings.ToUpper(entry.Level.String())),
		entry.Message)
	if entry.Data != nil && len(entry.Data) > 0 {
		msg = fmt.Sprintf("%s\n%s\n", msg, util.PrettyMarshal(entry.Data))
	}

	return []byte(msg), nil
}

func Init(level string) {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		logger.Fatal(err)
	}
	logger.SetFormatter(new(costumeLogFormatter))
	logger.SetLevel(lv)
	logger.Info(fmt.Sprintf("log level: %s", logrus.GetLevel().String()))
}
