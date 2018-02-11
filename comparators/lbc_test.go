package comparators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_hashString_success(t *testing.T) {
	require.Equal(t, "153fa238cec90e5a24b85a79109f91ebe68ca481", hashString("qwerty1234"))
}

func Test_randomString_success(t *testing.T) {
	salt := randomString(128)
	require.Equal(t, 128, len(salt))
}

func TestSetRandomSalt_success(t *testing.T) {
	p := &LBCPassword{}
	p.Plain = "a"
	p.Hash = "b"
	p.SetRandomSalt()
	assert.Equal(t, 16, len(p.Salt))
	assert.Equal(t, "", p.Plain)
	assert.Equal(t, "", p.Hash)
}

func TestSetPassword_success(t *testing.T) {
	p := &LBCPassword{}
	p.Salt = "d2rsph111lxo3twk"
	err := p.SetPlainPassword("qwerty1234")
	require.NoError(t, err)
	assert.Equal(t, p.Salt+"a829f192f7fd38700cacdc5c645596ce3e9d09b1", p.Hash)
	assert.Equal(t, "qwerty1234", p.Plain)
}

func TestSetPassword_without_salt(t *testing.T) {
	p := &LBCPassword{}
	err := p.SetPlainPassword("qwerty1234")
	require.Error(t, err)
}

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
