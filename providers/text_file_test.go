package providers

import (
	"testing"

	"os"

	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pb "gopkg.in/cheggaaa/pb.v1"
)

var testFile = "test.txt"

func createFile(t *testing.T) {
	f, err := os.Create(testFile)
	require.NoError(t, err)
	_, err = f.WriteString("azerty1234\nqwerty1234\nfoo\nbar\n")
	require.NoError(t, err)
}

func removeFile(t *testing.T) {
	os.Remove(testFile)
}

func TestPrepare_success(t *testing.T) {
	createFile(t)
	d := NewDictionaryFromFile(testFile)
	size, err := d.Prepare()
	require.NoError(t, err)
	assert.Equal(t, 30, size)
	removeFile(t)
}

func TestPrepare_error_open_file(t *testing.T) {
	removeFile(t)
	d := NewDictionaryFromFile(testFile)
	_, err := d.Prepare()
	require.Error(t, err)
}

func TestClose_error(t *testing.T) {
	d := &TextFileDictionary{}
	err := d.Close()
	require.Error(t, err)
}

func TestClose_success(t *testing.T) {
	createFile(t)
	d := NewDictionaryFromFile(testFile)
	_, err := d.Prepare()
	require.NoError(t, err)
	err = d.Close()
	require.NoError(t, err)
	removeFile(t)
}

func TestNext_success(t *testing.T) {
	createFile(t)
	d := NewDictionaryFromFile(testFile)

	_, err := d.Prepare()
	require.NoError(t, err)

	bar := &pb.ProgressBar{}
	d.SetProgressBar(bar.NewProxyReader(d.GetReader()))

	assert.True(t, d.Next())
	assert.Equal(t, "azerty1234", d.Value())

	assert.True(t, d.Next())
	assert.Equal(t, "qwerty1234", d.Value())

	assert.True(t, d.Next())
	assert.Equal(t, "foo", d.Value())

	assert.True(t, d.Next())
	assert.Equal(t, "bar", d.Value())

	//Last line
	assert.False(t, d.Next())

	removeFile(t)
}

func TestErr_error(t *testing.T) {
	d := &TextFileDictionary{}
	d.err = errors.New("foo")
	err := d.Err()
	require.Error(t, err)
}
