package cmd

import (
	"os"

	"github.com/pkg/errors"
	pb "gopkg.in/cheggaaa/pb.v1"
)

type DictionaryProvider interface {
	Prepare() (int, error)
	GetReader() *os.File
	SetProgressBar(reader *pb.Reader)
	Next() bool
	Value() string
	Err() error
	Close() error
}

func readDictionary(p DictionaryProvider, dictionaryChan chan string, crashChan chan error) {
	defer close(dictionaryChan)

	// Init provider
	size, err := p.Prepare()
	if err != nil {
		crashChan <- errors.Wrap(err, "unable to prepare dictionary provider")
	}

	// Progress bar
	bar := pb.New(size).SetUnits(pb.U_BYTES).Start()
	defer bar.Finish()
	p.SetProgressBar(bar.NewProxyReader(p.GetReader()))

	// Read values
	for p.Next() {
		dictionaryChan <- p.Value()
	}

	// Last provider error
	if p.Err() != nil {
		crashChan <- errors.Wrap(p.Err(), "dictionary provider error")
	}

	// Close provider
	err = p.Close()
	if err != nil {
		crashChan <- errors.Wrap(err, "unable to close dictionary provider")
	}
}
