package comparators

import (
	"sync"

	"github.com/jclebreton/hash-cracker/providers"
	"github.com/sirupsen/logrus"
	"runtime"
)

type PasswordComparator interface {
	GetHash() string
	SetHash(hash string)
	Compare(plain string) bool
}

const dictionaryBuffer = 10000

func Compare(comparator PasswordComparator, p providers.DictionaryProvider) {
	wg := sync.WaitGroup{}
	gracefulChan := make(chan struct{})
	crashChan := make(chan error)
	dictionaryChan := make(chan string, dictionaryBuffer)
	ResultChan := make(chan string)


	// Dictionary provider
	go providers.Read(p, dictionaryChan, crashChan, gracefulChan)

	// Init comparators workers
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(&wg, gracefulChan, dictionaryChan, ResultChan, comparator)
	}

	logrus.Infof("cracking hash: %s with %d logical CPUs and %d go routines", comparator.GetHash(),
		runtime.NumCPU(), runtime.NumGoroutine())

	// Provider error
	go func() {
		err := <-crashChan
		logrus.WithError(err).Error("dictionary provider error")
		close(gracefulChan)
	}()

	// Success
	go func() {
		plainPassword := <-ResultChan
		logrus.WithField("plain", plainPassword).Info("password found")
		close(gracefulChan)
	}()

	wg.Wait()
}

func worker(wg *sync.WaitGroup, gracefulChan chan struct{}, dictionaryChan chan string, ResultChan chan string, c PasswordComparator) {
	defer wg.Done()
	for {
		select {
		case passwd := <-dictionaryChan:
			if c.Compare(passwd) {
				ResultChan <- passwd
				return
			}
		case <-gracefulChan:
			return
		}
	}
}
