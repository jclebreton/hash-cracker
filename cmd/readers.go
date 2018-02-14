package cmd

import (
	"time"

	"github.com/jclebreton/hash-cracker/dictionaries"
	"github.com/jclebreton/hash-cracker/hashers"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	pb "gopkg.in/cheggaaa/pb.v1"
)

// DictionaryReader returns the dictionary file
func DictionaryReader(p dictionaries.Provider, n int) ([]string, error) {
	// Init provider
	if err := p.Prepare(); err != nil {
		return nil, errors.Wrap(err, "unable to prepare dictionary provider")
	}

	// Progress bar
	bar := pb.New(p.GetTotal()).SetUnits(pb.U_NO).Prefix(p.GetName()).Start()
	defer bar.Finish()

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
func HashesReader(p dictionaries.Provider, errChan chan error, hashesChans map[int]chan Hash, hasher hashers.Hasher) {
	// Init provider
	if err := p.Prepare(); err != nil {
		errChan <- errors.Wrap(err, "unable to prepare dictionary provider")
	}

	// Progress bar
	bar := pb.New(p.GetTotal()).SetUnits(pb.U_NO).Prefix(p.GetName()).Start()
	defer bar.Finish()

	// Read values and sent them to workers
	for p.Next() {

		//Build hash
		hash := p.Value()
		h := Hash{}
		h.SetHasher(hasher)
		if err := h.SetHash(hash); err != nil {
			logrus.WithField("hash", hash).WithError(err).Error("HashesReader error")
			continue
		}

		// Send the same hash to all workers
		for workerID, _ := range hashesChans {
			hashesChans[workerID] <- h
		}

		bar.Increment()
	}

	// Last provider error
	if p.Err() != nil {
		errChan <- errors.Wrap(p.Err(), "dictionary provider error")
	}

	// Close provider
	if err := p.Close(); err != nil {
		errChan <- errors.Wrap(err, "unable to close dictionary provider")
	}

	time.Sleep(time.Second)
	closeWorkers(hashesChans)
}
