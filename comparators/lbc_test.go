package comparators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetHash_success(t *testing.T) {
	p := &LBCPassword{}
	p.SetHash("d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1")
	assert.Equal(t, p.Salt+"a829f192f7fd38700cacdc5c645596ce3e9d09b1", p.Hash)
	assert.Equal(t, "d2rsph111lxo3twk", p.Salt)
	assert.Equal(t, "", p.Plain)
}

func TestCompare_success(t *testing.T) {
	p := &LBCPassword{}
	p.SetHash("d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1")
	assert.True(t, p.Compare("qwerty1234"))
	assert.False(t, p.Compare("azerty1234"))
}
