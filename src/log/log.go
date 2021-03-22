package log

import (
	log "github.com/sirupsen/logrus"
	"io"
)

var ClientLog *log.Logger

func SetUpClientLog(out io.Writer, level string) error {
	ClientLog = log.New()
	ClientLog.SetOutput(out)
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return err
	}
	ClientLog.SetLevel(lvl)
	return nil
}
