package main

import (
	"sync"

	"github.com/jclebreton/hash-cracker/cmd"
	"github.com/jclebreton/hash-cracker/comparators"
	"github.com/jclebreton/hash-cracker/providers"
	"github.com/sirupsen/logrus"
)

var concurrency = 200

func init() {
	logrus.Infof("using %d go routines", concurrency)
}

// Overridden at compile time when using script/build.sh
var version = "dev"
var buildDate = "no build date"

func main() {
	cmd.Execute()
	//path := "test.txt"
	path := "crackstation.txt"
	comparator := &comparators.LBCPassword{}
	comparator.SetRandomSalt()
	comparator.SetPlainPassword("qwerty1234")
	comparator.Plain = ""

	// Init
	wg := sync.WaitGroup{}
	gracefulChan := make(chan bool)
	crashChan := make(chan error)
	dictionaryChan := make(chan string, 100)
	ResultChan := make(chan string)

	logrus.Infof("cracking hash: %s", comparator.Hash)

	// Init dictionary provider
	p := providers.NewDictionaryFromFile(path)
	go readDictionary(p, dictionaryChan, crashChan)

	// Init comparators workers
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(&wg, gracefulChan, dictionaryChan, ResultChan, comparator)
	}

	// Provider error
	go func() {
		err := <-crashChan
		logrus.WithError(err).Error("dictionary provider error")
		graceful(gracefulChan)
	}()

	// Success
	go func() {
		plainPassword := <-ResultChan
		logrus.WithField("plain", plainPassword).Info("password found")
		graceful(gracefulChan)
	}()

	wg.Wait()
}
