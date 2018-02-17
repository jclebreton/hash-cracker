package cmd

import (
	"sync"

	"os"

	"github.com/jclebreton/hash-cracker/dictionaries"
	"github.com/jclebreton/hash-cracker/hashers"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/cheggaaa/pb.v1"
)

// Run will start the process
func Run(h dictionaries.Provider, d dictionaries.Provider, hasher hashers.Hasher, nbWorkers int) {
	logrus.Infof("%d workers", nbWorkers)

	wg := sync.WaitGroup{}
	errChan := make(chan error)
	resultChan := make(chan map[int]Hash)

	//Progression bars
	pb1 := pb.New(0).SetUnits(pb.U_NO).Prefix("Dictionary")
	pb2 := pb.New(0).SetUnits(pb.U_NO).Prefix("    Hashes")
	pb3 := pb.New(0).SetUnits(pb.U_NO).Prefix("   Cracked")
	pool, err := pb.StartPool(pb1, pb2, pb3)
	pb1.ShowPercent = true
	pb2.ShowPercent = true
	pb3.ShowPercent = true
	if err != nil {
		logrus.WithError(err).Fatal("progress bar error")
	}

	// Read dictionary
	dictionary, err := DictionaryReader(pb1, d)
	if err != nil {
		logrus.WithError(err).Error("dictionary provider error")
	}
	dictionaries := splitSlice(dictionary, nbWorkers)
	if len(dictionaries) < nbWorkers {
		nbWorkers = len(dictionaries)
		logrus.Infof("Reduce the number of workers to %d", nbWorkers)
	}

	// Init workers
	hashesChans := make(map[int]chan Hash, 10000)
	resetChans := make(map[int]chan struct{})
	for i := 1; i <= nbWorkers; i++ {
		wg.Add(1)
		resetChans[i], hashesChans[i] = worker(i, &wg, errChan, dictionaries[i-1], resultChan)
	}

	// Provider error
	go func() {
		f, err := os.OpenFile("error.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.WithError(err).Fatal("unable to open error file")
		}
		defer f.Close()

		err = <-errChan

		if _, err = f.WriteString(err.Error()); err != nil {
			logrus.WithError(err).Fatal("unable to save error")
		}

		logrus.WithError(err).Error("error")
		for k, _ := range hashesChans {
			close(hashesChans[k])
		}
	}()

	// Read hashes
	go HashesReader(pb2, pb3, h, errChan, hashesChans, hasher)

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
				pb3.Increment()

				text := hash.GetHash() + "\t" + hash.GetPlain() + "\n"
				if _, err = f.WriteString(text); err != nil {
					logrus.WithError(err).Fatal("unable to save results")
				}

				//Reset all worker excepted
				for k, _ := range resetChans {
					if k != chanID {
						resetChans[k] <- struct{}{}
					}
				}
			}
		}
	}()

	wg.Wait()
	pool.Stop()
}

func worker(id int, wg *sync.WaitGroup, errChan chan error, dictionary []string, resultChan chan map[int]Hash) (chan struct{}, chan Hash) {
	resetChan := make(chan struct{})
	hashesChan := make(chan Hash)

	go func() {
		defer wg.Done()
		for hash := range hashesChan {
		start:
			for _, plain := range dictionary {
				select {
				case <-resetChan:
					logrus.WithField("worker", id).WithField("hash", hash.GetHash()).Debug("reset")
					break start
				default:
					var ok bool
					var err error
					if ok, err = hash.Compare(plain); err != nil {
						logrus.WithError(err).Error("unable to compare hash")
						errChan <- errors.Wrap(err, "unable to compare hash")
						return
					}
					if ok {
						hash.SetPlain(plain)
						resultChan <- map[int]Hash{id: hash}
						logrus.WithField("worker", id).WithField("hash", hash.GetHash()).Debug("found")
						break start
					}
				}
			}
		}
	}()

	return resetChan, hashesChan
}

func closeWorkers(hashesChans map[int]chan Hash) {
	for workerID, _ := range hashesChans {
		close(hashesChans[workerID])
	}
}
func splitSlice(s []string, n int) [][]string {
	var divided [][]string
	chunkSize := (len(s) + n - 1) / n
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		divided = append(divided, s[i:end])
	}

	return divided
}
