package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash_properties(t *testing.T) {
	hash := Hash{
		Hash:  "a",
		Salt:  "b",
		Plain: "c",
	}
	assert.Equal(t, "a", hash.Hash)
	assert.Equal(t, "b", hash.Salt)
	assert.Equal(t, "c", hash.Plain)
}
