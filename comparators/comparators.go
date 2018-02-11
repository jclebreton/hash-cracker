package comparators

import (
	"sync"

	"github.com/jclebreton/hash-cracker/providers"
	"github.com/sirupsen/logrus"
)

type PasswordComparator interface {
	GetHash() string
	SetHash(hash string)
	Compare(plain string) bool
}

var concurrency = 200

func Compare(comparator PasswordComparator, p providers.DictionaryProvider) {
	wg := sync.WaitGroup{}
	gracefulChan := make(chan bool)
	crashChan := make(chan error)
	dictionaryChan := make(chan string)
	ResultChan := make(chan string)

	logrus.Infof("cracking hash: %s", comparator.GetHash())

	// Dictionary provider
	go providers.Read(p, dictionaryChan, crashChan)

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

func worker(wg *sync.WaitGroup, gracefulChan chan bool, dictionaryChan chan string, ResultChan chan string, c PasswordComparator) {
	for {
		select {
		case passwd := <-dictionaryChan:
			if c.Compare(passwd) {
				ResultChan <- passwd
				wg.Done()
				return
			}
		case <-gracefulChan:
			wg.Done()
			return
		}
	}
}

func graceful(gracefulChan chan bool) {
	for i := 0; i < concurrency; i++ {
		gracefulChan <- true
	}
}
