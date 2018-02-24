package usecases

import (
	"sync"

	"os"

	"strings"

	"fmt"

	"time"

	"github.com/jclebreton/hash-cracker/domains"
	"github.com/jclebreton/hash-cracker/infrastructures/comparators"
	"github.com/jclebreton/hash-cracker/infrastructures/progress"
	"github.com/jclebreton/hash-cracker/infrastructures/readers"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var startDate time.Time

func init() {
	startDate = time.Now()
}

// CrackHashesUsingDictionaryHandler is the interactor
type CrackHashesUsingDictionaryHandler struct {
	HashComparator    comparators.Comparator
	DictionaryReader  readers.DictionaryReader
	HashesReader      readers.HashesReader
	ProgressBarPooler progress.BarPooler
}

// Handle cracks hashes using dictionary
func (handler *CrackHashesUsingDictionaryHandler) Handle(nbWorkers int, randomize bool) {
	wg := sync.WaitGroup{}
	errChan := make(chan error)
	resultChan := make(chan map[int]domains.Hash, 10000)

	// Read dictionary
	dictionaries, err := handler.DictionaryReader.Reader(randomize, nbWorkers)
	if err != nil {
		logrus.WithError(err).Error("dictionary provider error")
	}
	if len(dictionaries) < nbWorkers {
		nbWorkers = len(dictionaries)
	}
	logrus.Infof("%d workers", nbWorkers)

	// Init workers
	hashesChans := make(map[int]chan domains.Hash, 10000)
	resetChans := make(map[int]chan domains.Hash)
	var progressBar *progress.CheggaaBar
	for i := 1; i <= nbWorkers; i++ {
		wg.Add(1)
		resetChans[i], hashesChans[i], progressBar = worker(i, &wg, errChan, dictionaries[i-1], resultChan, handler.HashComparator)
		handler.ProgressBarPooler.Add(progressBar)
		logrus.WithField("id", i).Debug("worker started")
	}

	handler.ProgressBarPooler.Add(handler.HashesReader.ProgressBarHashes)
	handler.ProgressBarPooler.Add(handler.HashesReader.ProgressBarCracked)

	err = handler.ProgressBarPooler.Start()
	if err != nil {
		logrus.WithError(err).Fatal("progress bar error")
	}

	// Provider error
	go func() {
		f, err := os.OpenFile("error.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.WithError(err).Fatal("unable to open error file")
		}
		defer f.Close()

		logrus.Debug("error routine started")

		err = <-errChan
		logrus.WithError(err).Error("error catched")
		if _, err = f.WriteString(err.Error()); err != nil {
			logrus.WithError(err).Fatal("unable to save error")
		}
		for k := range hashesChans {
			logrus.WithField("worker", k).Debug("trying to close workers")
			close(hashesChans[k])
		}
		logrus.Debug("workers closed")
	}()

	// Read hashes
	go handler.HashesReader.Reader(errChan, hashesChans)

	// Success
	go func() {
		f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.WithError(err).Fatal("unable to open output file")
		}
		defer f.Close()

		for {
			result := <-resultChan
			for chanID, hash := range result {
				text := hash.Hash + "\t" + hash.Plain + "\n"
				if _, err = f.WriteString(text); err != nil {
					logrus.WithError(err).Fatal("unable to save results")
				}

				//Reset all worker excepted
				for k := range resetChans {
					if k != chanID {
						resetChans[k] <- hash
					}
				}

				handler.HashesReader.ProgressBarCracked.Increment()
			}
		}
	}()

	wg.Wait()
	handler.ProgressBarPooler.Stop()

	logrus.Infof("finish in %s", time.Now().Sub(startDate))
}

func worker(id int, wg *sync.WaitGroup, errChan chan error, dictionary []string, resultChan chan map[int]domains.Hash,
	hashComparator comparators.Comparator) (chan domains.Hash, chan domains.Hash, *progress.CheggaaBar) {

	logrus.WithField("Worker", id).WithField("words", dictionary).Debug("Dictionary")
	resetChan := make(chan domains.Hash, 100)
	hashesChan := make(chan domains.Hash)
	progressBar := progress.NewProgressBar(fmt.Sprintf("worker %d", id))
	progressBar.SetTotal(int64(len(dictionary)))

	go func() {
		defer wg.Done()
		for hash := range hashesChan {
			progressBar.Set(0)
		start:
			for _, plain := range dictionary {
				select {
				case hashToReset := <-resetChan:
					if strings.Compare(hashToReset.Hash, hash.Hash) == 0 {
						logrus.WithField("worker", id).WithField("hash", hash.Hash).Debug("reset")
						break start
					}
				default:
					ok, err := hashComparator.Compare(hash, plain)
					if err != nil {
						logrus.WithError(err).Error("unable to compare hash")
						errChan <- errors.Wrap(err, "unable to compare hash")
						return
					} else if ok {
						hash.Plain = plain
						resultChan <- map[int]domains.Hash{id: hash}
						logrus.WithField("worker", id).WithField("hash", hash.Hash).Debug("found")
						break start
					}
				}
				progressBar.Increment()
			}
		}
	}()

	return resetChan, hashesChan, progressBar
}
