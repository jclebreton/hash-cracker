package readers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitSlice_nbworkers(t *testing.T) {
	tests := []struct {
		dictionary []string

		nbWorkers     int
		expectedNbDic int
	}{
		{[]string{"a", "b", "c", "d", "e", "f", "g", "h"}, 4, 4},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, 4, 4},
		{[]string{"a", "b", "c", "d", "e", "f"}, 4, 4},
		{[]string{"a", "b", "c", "d", "e"}, 4, 4},
		{[]string{"a", "b", "c", "d"}, 4, 4},
		{[]string{"a", "b", "c"}, 4, 3},
		{[]string{"a", "b"}, 4, 2},
		{[]string{"a"}, 4, 1},
		{[]string{}, 4, 0},
	}
	for _, tt := range tests {
		dictionaries := splitSlice(tt.dictionary, tt.nbWorkers)
		require.Equal(t, tt.expectedNbDic, len(dictionaries))
	}
}

func Test_splitSlice_content(t *testing.T) {
	dictionary := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	dictionaries := splitSlice(dictionary, 4)
	require.Equal(t, 4, len(dictionaries))

	require.Equal(t, 3, len(dictionaries[0]))
	require.Equal(t, "a", dictionaries[0][0])
	require.Equal(t, "b", dictionaries[0][1])
	require.Equal(t, "c", dictionaries[0][2])

	require.Equal(t, 2, len(dictionaries[1]))
	require.Equal(t, "d", dictionaries[1][0])
	require.Equal(t, "e", dictionaries[1][1])

	require.Equal(t, 2, len(dictionaries[2]))
	require.Equal(t, "f", dictionaries[2][0])
	require.Equal(t, "g", dictionaries[2][1])

	require.Equal(t, 2, len(dictionaries[3]))
	require.Equal(t, "h", dictionaries[3][0])
	require.Equal(t, "i", dictionaries[3][1])
}
