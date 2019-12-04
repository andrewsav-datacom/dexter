package main

import (
	"github.com/gini/dexter/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/mattn/go-colorable"
)

func main() {
	// set log format & level
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true,ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
	log.SetLevel(log.InfoLevel)

	cmd.Execute()
}
