package cmd

import (
	"time"

	"fmt"

	"github.com/jclebreton/hash-cracker/dictionaries"
	"github.com/jclebreton/hash-cracker/hashers"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	pb "gopkg.in/cheggaaa/pb.v1"
)

// DictionaryReader returns the dictionary file
func DictionaryReader(bar *pb.ProgressBar, p dictionaries.Provider) ([]string, error) {
	defer bar.Finish()

	// Init provider
	if err := p.Prepare(); err != nil {
		return nil, errors.Wrap(err, "unable to prepare dictionary provider")
	}

	bar.Total = p.GetTotal()

	// Read values
	result := []string{}
	for p.Next() {
		result = append(result, p.Value())
		bar.Increment()
	}

	// Last provider error
	if p.Err() != nil {
		return nil, errors.Wrap(p.Err(), "dictionary provider error")
	}

	// Close provider
	if err := p.Close(); err != nil {
		return nil, errors.Wrap(err, "unable to close dictionary provider")
	}

	return result, nil
}

// HashesReader returns the hashes
func HashesReader(bar1 *pb.ProgressBar, bar2 *pb.ProgressBar, p dictionaries.Provider, errChan chan error, hashesChans map[int]chan Hash, hasher hashers.Hasher) {
	defer bar1.Finish()

	// Init provider
	if err := p.Prepare(); err != nil {
		errChan <- errors.Wrap(err, "unable to prepare dictionary provider")
		return
	}

	bar1.Total = p.GetTotal()
	var current int64

	// Read values and sent them to workers
	for p.Next() {

		//Build hash
		hash := Hash{}
		hash.SetHasher(hasher)
		if err := hash.SetHash(p.Value()); err != nil {
			logrus.WithField("hash", hash.GetHash()).WithError(err).Error("HashesReader error")
			errChan <- errors.Wrap(err, fmt.Sprintf("hash (%s) error", hash.GetHash()))
			return
		}

		// Send the same hash to all workers
		for workerID, _ := range hashesChans {
			hashesChans[workerID] <- hash
		}

		bar1.Increment()
		current++
		bar2.Total = current
	}

	// Last provider error
	if p.Err() != nil {
		errChan <- errors.Wrap(p.Err(), "dictionary provider error")
		return
	}

	// Close provider
	if err := p.Close(); err != nil {
		errChan <- errors.Wrap(err, "unable to close dictionary provider")
		return
	}

	time.Sleep(time.Second)
	closeWorkers(hashesChans)
}
