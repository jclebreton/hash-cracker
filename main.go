package main

import (
	"github.com/jclebreton/hash-cracker/cli"
	"github.com/jclebreton/hash-cracker/infrastructures/comparators"
	"github.com/sirupsen/logrus"
)

// Overridden at compile time when using script/build.sh
var version = "dev"
var buildDate = "no build date"

func main() {
	hasher := &comparators.Sha1WithSalt{}
	if err := cli.InitRootCmd(hasher).Execute(); err != nil {
		logrus.WithError(err).Fatal("Something wrong happens when running the command.")
	}
}
