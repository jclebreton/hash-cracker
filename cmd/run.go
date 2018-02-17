package cmd

import (
	"sync"

	"os"

	"strings"

	"fmt"

	"github.com/jclebreton/hash-cracker/dictionaries"
	"github.com/jclebreton/hash-cracker/hashers"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/cheggaaa/pb.v1"
)

// Run will start the process
func Run(h dictionaries.Provider, d dictionaries.Provider, hasher hashers.Hasher, nbWorkers int) {
	wg := sync.WaitGroup{}
	errChan := make(chan error)
	resultChan := make(chan map[int]Hash, 10000)

	// Read dictionary
	pbDictionary := pb.New(0).SetUnits(pb.U_NO).Prefix("Dictionary")
	pbDictionary.Start()
	pbDictionary.ShowPercent = true
	pbDictionary.ShowTimeLeft = true
	dictionary, err := DictionaryReader(pbDictionary, d)
	if err != nil {
		logrus.WithError(err).Error("dictionary provider error")
	}
	dictionaries := splitSlice(dictionary, nbWorkers)
	if len(dictionaries) < nbWorkers {
		nbWorkers = len(dictionaries)
	}
	logrus.Infof("%d workers", nbWorkers)

	// Init workers
	hashesChans := make(map[int]chan Hash, 10000)
	resetChans := make(map[int]chan Hash)
	var progressBar *pb.ProgressBar
	var progressBars []*pb.ProgressBar

	for i := 1; i <= nbWorkers; i++ {
		wg.Add(1)
		resetChans[i], hashesChans[i], progressBar = worker(i, &wg, errChan, dictionaries[i-1], resultChan)
		progressBars = append(progressBars, progressBar)
		logrus.WithField("id", i).Debug("worker started")
	}

	pbHashes := pb.New(0).SetUnits(pb.U_NO).Prefix("Hashes")
	progressBars = append(progressBars, pbHashes)
	pbCracked := pb.New(0).SetUnits(pb.U_NO).Prefix("Cracked")
	progressBars = append(progressBars, pbCracked)
	pool, err := pb.StartPool(progressBars...)
	pbHashes.ShowPercent = true
	pbHashes.ShowTimeLeft = true
	pbCracked.ShowPercent = true
	pbCracked.ShowTimeLeft = false

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
		for k, _ := range hashesChans {
			logrus.WithField("worker", k).Debug("trying to close workers")
			close(hashesChans[k])
		}
		logrus.Debug("workers closed")
	}()

	// Read hashes
	go HashesReader(pbHashes, pbCracked, h, errChan, hashesChans, hasher)

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
				pbCracked.Increment()

				text := hash.GetHash() + "\t" + hash.GetPlain() + "\n"
				if _, err = f.WriteString(text); err != nil {
					logrus.WithError(err).Fatal("unable to save results")
				}

				//Reset all worker excepted
				for k, _ := range resetChans {
					if k != chanID {
						resetChans[k] <- hash
					}
				}
			}
		}
	}()

	wg.Wait()
	pool.Stop()
}

func worker(id int, wg *sync.WaitGroup, errChan chan error, dictionary []string, resultChan chan map[int]Hash) (chan Hash, chan Hash, *pb.ProgressBar) {
	logrus.WithField("Worker", id).WithField("words", dictionary).Debug("Dictionary")
	resetChan := make(chan Hash)
	hashesChan := make(chan Hash)
	bar := pb.New(len(dictionary)).SetUnits(pb.U_NO).Prefix(fmt.Sprintf("worker %d", id))
	bar.ShowPercent = true
	bar.ShowTimeLeft = false

	go func() {
		defer wg.Done()
		for hash := range hashesChan {
			bar.Set(0)
		start:
			for _, plain := range dictionary {
				select {
				case hashToReset := <-resetChan:
					if strings.Compare(hashToReset.GetHash(), hash.GetHash()) == 0 {
						logrus.WithField("worker", id).WithField("hash", hash.GetHash()).Debug("reset")
						break start
					}
				default:
					ok, err := hash.Compare(plain)
					if err != nil {
						logrus.WithError(err).Error("unable to compare hash")
						errChan <- errors.Wrap(err, "unable to compare hash")
						return
					} else if ok {
						hash.SetPlain(plain)
						resultChan <- map[int]Hash{id: hash}
						logrus.WithField("worker", id).WithField("hash", hash.GetHash()).Debug("found")
						break start
					}
				}
				bar.Increment()
			}
		}
	}()

	return resetChan, hashesChan, bar
}

func closeWorkers(hashesChans map[int]chan Hash) {
	for workerID, _ := range hashesChans {
		close(hashesChans[workerID])
	}
}
func splitSlice(s []string, n int) [][]string {
	var divided [][]string
	var quotient, remainder int
	var end int

	quotient = len(s) / n
	remainder = len(s) % n

	for i := 0; i < n; i++ {
		start := end
		end = start + quotient

		if remainder > 0 {
			end++
			remainder--
		}

		if end > len(s) {
			end = len(s)
		}

		if len(s[start:end]) != 0 {
			divided = append(divided, s[start:end])
		}
	}

	return divided
}
