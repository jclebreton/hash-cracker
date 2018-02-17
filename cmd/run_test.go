package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseResponse_invalid(t *testing.T) {
	tests := []struct {
		dictionary      []string
		nbWorkers       int
		expectedNbDic   int
		expectedNbWords int
	}{
		{[]string{"a", "b", "c", "d", "e", "f", "g", "h"}, 4, 4, 2},
		{[]string{"a", "b", "c", "d"}, 4, 4, 1},
		{[]string{"a", "b"}, 4, 2, 1},
		{[]string{"a"}, 4, 1, 1},
		{[]string{}, 4, 0, 0},
	}
	for _, tt := range tests {
		dictionaries := splitSlice(tt.dictionary, tt.nbWorkers)
		require.Equal(t, tt.expectedNbDic, len(dictionaries))
		index := 0
		for k, _ := range dictionaries {
			assert.Equal(t, tt.expectedNbWords, len(dictionaries[k]))
			for j, word := range dictionaries[k] {
				assert.Equal(t, tt.dictionary[index+j], word)
			}
			index += len(dictionaries[k])
		}
	}
}

func Test_splitSlice_with_dictionaries_modulo(t *testing.T) {
	dictionary := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	dictionaries := splitSlice(dictionary, 4)
	require.Equal(t, 4, len(dictionaries))
	assert.Equal(t, 3, len(dictionaries[0]))
	assert.Equal(t, 3, len(dictionaries[1]))
	assert.Equal(t, 3, len(dictionaries[2]))
	assert.Equal(t, 1, len(dictionaries[3]))
}
