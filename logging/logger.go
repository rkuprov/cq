package l

import "log"

var Enabled *bool

type logger struct {
	logger *log.Logger
}

var l logger

func New(logger *log.Logger) {
	l.logger = logger
}

func Debug(msg string) {
	if *Enabled {
		l.logger.Println(msg)
	}
}
