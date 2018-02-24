package readers

import (
	"testing"

	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrepare_success(t *testing.T) {
	createFile(t)
	d := NewTextFileReader(testFile)
	err := d.Prepare()
	require.NoError(t, err)
	assert.Equal(t, int64(4), d.GetTotal())
	removeFile(t)
}

func TestGetName_success(t *testing.T) {
	createFile(t)
	d := NewTextFileReader(testFile)
	err := d.Prepare()
	require.NoError(t, err)
	assert.Equal(t, testFile, d.GetName())
	removeFile(t)
}

func TestPrepare_error_open_file(t *testing.T) {
	removeFile(t)
	d := NewTextFileReader(testFile)
	err := d.Prepare()
	require.Error(t, err)
}

func TestClose_error(t *testing.T) {
	d := &TextFileReader{}
	err := d.Close()
	require.Error(t, err)
}

func TestClose_success(t *testing.T) {
	createFile(t)
	d := NewTextFileReader(testFile)
	err := d.Prepare()
	require.NoError(t, err)
	err = d.Close()
	require.NoError(t, err)
	removeFile(t)
}

func TestNext_success(t *testing.T) {
	createFile(t)
	d := NewTextFileReader(testFile)

	err := d.Prepare()
	require.NoError(t, err)

	assert.True(t, d.Next())
	assert.Equal(t, "azerty1234", d.Value())

	assert.True(t, d.Next())
	assert.Equal(t, "qwerty1234", d.Value())

	assert.True(t, d.Next())
	assert.Equal(t, "foo", d.Value())

	assert.True(t, d.Next())
	assert.Equal(t, "Bar", d.Value())

	//Last line
	assert.False(t, d.Next())

	removeFile(t)
}

func TestErr_error(t *testing.T) {
	d := &TextFileReader{}
	d.err = errors.New("foo")
	err := d.Err()
	require.Error(t, err)
}
