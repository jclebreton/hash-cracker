package main

import (
	"github.com/jclebreton/hash-cracker/cmd"
	"github.com/sirupsen/logrus"
)

// Overridden at compile time when using script/build.sh
var version = "dev"
var buildDate = "no build date"

func main() {
	if err := cmd.InitRootCmd().Execute(); err != nil {
		logrus.WithError(err).Fatal("Something wrong happen when running the command.")
	}
}
