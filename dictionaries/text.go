package dictionaries

import (
	"bufio"
	"os"

	"bytes"
	"io"

	"github.com/pkg/errors"
)

// TextFile contains all file meta data
type TextFile struct {
	path    string
	scanner *bufio.Scanner
	value   string
	err     error
	file    *os.File

	total   int
	current int
}

// NewDictionaryFromFile is the constructor
func New(path string) *TextFile {
	return &TextFile{path: path}
}

// GetName returns the provider name
func (d *TextFile) GetName() string {
	return d.path
}

// Prepare initializes the dictionary source
func (d *TextFile) Prepare() error {
	//lines
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
func (d *TextFile) Next() bool {
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
func (d *TextFile) Value() string {
	return d.value
}

// Error returns the last error
func (d *TextFile) Err() error {
	return d.err
}

// Close closes the file
func (d *TextFile) Close() error {
	return d.file.Close()
}

// GetTotal returns the number of lines
func (d *TextFile) GetTotal() int {
	return d.total
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
