package comparators

import (
	"testing"

	"github.com/jclebreton/hash-cracker/domains"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompare_no_errors(t *testing.T) {
	tests := []struct {
		hash           string
		plain          string
		expectedResult bool
	}{
		{"d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1", "qwerty1234", true},
		{"d2rsph111lxo3twk39e169d94697bc5fc3e9da8bd17b0c23677a7583", "12345xxx", true},
		{"d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1", "a", false},
		{"d2rsph111lxo3twk39e169d94697bc5fc3e9da8bd17b0c23677a7583", "b", false},
	}
	for _, tt := range tests {
		hash := domains.Hash{Hash: tt.hash}
		comparator := &Sha1WithSalt{}
		result, err := comparator.Compare(hash, tt.plain)
		require.NoError(t, err)
		assert.Equal(t, tt.expectedResult, result)
	}
}

func TestCompare_errors(t *testing.T) {
	tests := []struct {
		hash           string
		plain          string
		expectedResult bool
	}{
		{"a829f192f7fd38700cacdc5c645596ce3e9d09b1", "qwerty1234", false},
		{"foo", "12345xxx", false},
		{"", "a", false},
		{"d2rsph111lxo3twk39e169d94697bc5fc3e9da8bd17b0c2da8bd17b0c23677a7583", "b", false},
	}
	for _, tt := range tests {
		hash := domains.Hash{Hash: tt.hash}
		comparator := &Sha1WithSalt{}
		_, err := comparator.Compare(hash, tt.plain)
		require.Error(t, err)
	}
}
