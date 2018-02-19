package readers

import (
	"github.com/jclebreton/hash-cracker/domains"
	"github.com/jclebreton/hash-cracker/infrastructures/progress"
	"github.com/pkg/errors"
)

type HashesReader struct {
	ProgressBarHashes  progress.ProgressBarer
	ProgressBarCracked progress.ProgressBarer
	HashesProvider     HashesProvider
}

// HashesReader returns the hashes
func (h *HashesReader) Reader(errChan chan error, hashesChans map[int]chan domains.Hash) {

	defer h.ProgressBarHashes.Finish()

	// Init provider
	if err := h.HashesProvider.Prepare(); err != nil {
		errChan <- errors.Wrap(err, "unable to prepare dictionary provider")
		return
	}

	h.ProgressBarHashes.SetTotal(h.HashesProvider.GetTotal())
	var current int64

	// Read values and sent them to workers
	for h.HashesProvider.Next() {
		hash := domains.Hash{Hash: h.HashesProvider.Value()}

		// Send the same hash to all workers
		for workerID, _ := range hashesChans {
			hashesChans[workerID] <- hash
		}

		h.ProgressBarHashes.Increment()
		current++
		h.ProgressBarCracked.SetTotal(current)
	}

	// Last provider error
	if h.HashesProvider.Err() != nil {
		errChan <- errors.Wrap(h.HashesProvider.Err(), "dictionary provider error")
		return
	}

	// Close provider
	if err := h.HashesProvider.Close(); err != nil {
		errChan <- errors.Wrap(err, "unable to close dictionary provider")
		return
	}

	closeWorkers(hashesChans)
}

func closeWorkers(hashesChans map[int]chan domains.Hash) {
	for workerID, _ := range hashesChans {
		close(hashesChans[workerID])
	}
}
