package providers

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/cheggaaa/pb.v1"
)

type TextFileDictionary struct {
	path    string
	scanner *bufio.Scanner
	value   string
	err     error
	file    *os.File
}

// NewDictionaryFromFile is the constructor
func NewDictionaryFromFile(path string) *TextFileDictionary {
	return &TextFileDictionary{path: path}
}

// Prepare initializes the dictionary source
func (d *TextFileDictionary) Prepare() (int, error) {
	file, err := os.OpenFile(d.path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return 0, errors.Wrap(err, "os.OpenFile error")
	}
	d.file = file

	//Size
	stat, err := d.file.Stat()
	if err != nil {
		return 0, errors.Wrap(err, "file.Stat error")
	}
	return int(stat.Size()), nil
}

// Next prepares the next value
func (d *TextFileDictionary) Next() bool {
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
func (d *TextFileDictionary) Value() string {
	return d.value
}

// Error returns the last error
func (d *TextFileDictionary) Err() error {
	return d.err
}

// Close closes the file
func (d *TextFileDictionary) Close() error {
	return d.file.Close()
}

// GetReader returns the file reader
func (d *TextFileDictionary) GetReader() *os.File {
	return d.file
}

// SetProgressBar overrides the reader by a proxy progress bar reader
func (d *TextFileDictionary) SetProgressBar(reader *pb.Reader) {
	d.scanner = bufio.NewScanner(reader)
}
