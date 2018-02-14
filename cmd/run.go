package cmd

import (
	"sync"

	"runtime"

	"os"

	"github.com/jclebreton/hash-cracker/dictionaries"
	"github.com/jclebreton/hash-cracker/hashers"
	"github.com/sirupsen/logrus"
	"gopkg.in/cheggaaa/pb.v1"
)

// Run will start the process
func Run(h dictionaries.Provider, d dictionaries.Provider, hasher hashers.Hasher) {
	wg := sync.WaitGroup{}
	errChan := make(chan error)
	resultChan := make(chan map[int]Hash)

	logrus.Infof("%d workers", runtime.NumCPU())

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
	dictionary, err := DictionaryReader(pb1, d, runtime.NumCPU())
	if err != nil {
		logrus.WithError(err).Error("dictionary provider error")
	}
	dic := splitSlice(dictionary, runtime.NumCPU())

	// Init workers
	hashesChans := make(map[int]chan Hash)
	resetChans := make(map[int]chan struct{})
	for i := 1; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		resetChans[i], hashesChans[i] = worker(i, &wg, dic[i-1], resultChan)
	}

	// Provider error
	go func() {
		err := <-errChan
		logrus.WithError(err).Error("dictionary provider error")
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
			logrus.WithError(err).Fatal("output file error")
		}
		defer f.Close()

		for {
			result := <-resultChan
			for chanID, hash := range result {
				pb3.Increment()

				text := hash.GetHash() + ":" + hash.GetPlain() + "\n"
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

func worker(id int, wg *sync.WaitGroup, dictionary []string, resultChan chan map[int]Hash) (chan struct{}, chan Hash) {
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
					if hash.Compare(plain) {
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
