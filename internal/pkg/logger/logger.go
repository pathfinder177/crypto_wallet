package logger

import (
	"io"
	"log"
	"os"
)

func New(logsPath string) (*log.Logger, *os.File) {
	f, err := os.OpenFile(logsPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Can not open %s to create logger\n", logsPath)
	}

	mw := io.MultiWriter(f, os.Stdout)
	l := log.New(mw, "", log.LstdFlags|log.LUTC)
	l.SetOutput(mw)

	return l, f
}
