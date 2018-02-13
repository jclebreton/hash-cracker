package comparators

import (
	"sync"

	"runtime"

	"github.com/jclebreton/hash-cracker/dictionaries"
	"github.com/sirupsen/logrus"
)

const dictionaryBuffer = 10000

// PasswordComparator is the interface used by comparator
type PasswordComparator interface {
	Compare(plain string) bool
	GetHash() string
}

// Compare will start the process
func Compare(comparator PasswordComparator, p dictionaries.DictionaryProvider) {
	wg := sync.WaitGroup{}
	gracefulChan := make(chan struct{})
	crashChan := make(chan error)
	dictionaryChan := make(chan string, dictionaryBuffer)
	ResultChan := make(chan string)

	// Dictionary provider
	go dictionaries.Read(p, dictionaryChan, crashChan, gracefulChan)

	// Init comparators workers
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(&wg, gracefulChan, dictionaryChan, ResultChan, comparator)
	}

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

	logrus.Infof("cracking hash: %s using %d workers", comparator.GetHash(), runtime.NumCPU())

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
