package readers

import (
	"github.com/jclebreton/hash-cracker/infrastructures/generators"
	"github.com/jclebreton/hash-cracker/infrastructures/progress"
	"github.com/pkg/errors"
)

type DictionaryReader struct {
	ProgressBar        progress.ProgressBarer
	DictionaryProvider DictionaryProvider
	PasswordsGenerator generators.Generator
}

// DictionaryReader returns the dictionary content
func (d *DictionaryReader) Reader(randomize bool, nbWorkers int) ([][]string, error) {
	defer d.ProgressBar.Finish()

	d.ProgressBar.Start()

	// Init provider
	if err := d.DictionaryProvider.Prepare(); err != nil {
		return nil, errors.Wrap(err, "unable to prepare dictionary provider")
	}

	d.ProgressBar.SetTotal(d.DictionaryProvider.GetTotal())

	// Read values
	result := []string{}
	for d.DictionaryProvider.Next() {
		if randomize {
			plains := d.PasswordsGenerator.Generate(d.DictionaryProvider.Value())
			result = append(result, plains...)
			n := len(plains)
			d.ProgressBar.IncrementTotal(int64(n - 1))
			d.ProgressBar.Add(int64(n))
		} else {
			result = append(result, d.DictionaryProvider.Value())
			d.ProgressBar.Increment()
		}
	}

	// Last provider error
	if d.DictionaryProvider.Err() != nil {
		return nil, errors.Wrap(d.DictionaryProvider.Err(), "dictionary provider error")
	}

	// Close provider
	if err := d.DictionaryProvider.Close(); err != nil {
		return nil, errors.Wrap(err, "unable to close dictionary provider")
	}

	return splitSlice(result, nbWorkers), nil
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
