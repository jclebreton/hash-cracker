package readers

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var testFile = "test.txt"

func createFile(t *testing.T) {
	f, err := os.Create(testFile)
	require.NoError(t, err)
	_, err = f.WriteString("azerty1234\nqwerty1234\nfoo\nBar\n")
	require.NoError(t, err)
}

func removeFile(t *testing.T) {
	os.Remove(testFile)
}
