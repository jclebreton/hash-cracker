package readers

import (
	"bufio"
	"os"

	"bytes"
	"io"

	"github.com/pkg/errors"
)

// TextFile contains all file meta data
type TextFileReader struct {
	path    string
	scanner *bufio.Scanner
	value   string
	err     error
	file    *os.File
	total   int64
}

// NewDictionaryFromFile is the constructor
func NewTextFileReader(path string) *TextFileReader {
	return &TextFileReader{path: path}
}

// GetName returns the provider name
func (d *TextFileReader) GetName() string {
	return d.path
}

// Prepare initializes the dictionary source
func (d *TextFileReader) Prepare() error {
	//count lines
	file, err := os.OpenFile(d.path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "os.OpenFile error")
	}
	if d.total, err = lineCounter(file); err != nil {
		return errors.Wrap(err, "os.OpenFile error")
	}

	//Open
	file, err = os.OpenFile(d.path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "os.OpenFile error")
	}
	d.file = file
	d.scanner = bufio.NewScanner(file)

	return nil
}

// Next prepares the next value
func (d *TextFileReader) Next() bool {
	if !d.scanner.Scan() {
		return false
	}

	d.value = d.scanner.Text()
	if d.scanner.Err() != nil {
		d.err = errors.Wrap(d.scanner.Err(), "scanner error")
		return false
	}

	return true
}

// Value returns the current value
func (d *TextFileReader) Value() string {
	return d.value
}

// Error returns the last error
func (d *TextFileReader) Err() error {
	return d.err
}

// Close closes the file
func (d *TextFileReader) Close() error {
	return d.file.Close()
}

// GetTotal returns the number of lines
func (d *TextFileReader) GetTotal() int64 {
	return d.total
}

func lineCounter(r io.Reader) (int64, error) {
	buf := make([]byte, 32*1024)
	var count int64
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += int64(bytes.Count(buf[:c], lineSep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
