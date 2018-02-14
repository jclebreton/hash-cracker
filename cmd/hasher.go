package cmd

import (
	"strings"

	"github.com/jclebreton/hash-cracker/hashers"
)

// Hash is the struct to store data
type Hash struct {
	hasher hashers.Hasher
	hash   string
	plain  string
}

// SetHasher is the hasher setter
func (p *Hash) SetHasher(hasher hashers.Hasher) {
	p.hasher = hasher
}

// GetHash is the hash getter
func (p *Hash) GetHash() string {
	return p.hash
}

// SetHash is the hash and salt setter
func (p *Hash) SetHash(hash string) error {
	if err := p.hasher.SetSaltFromHash(hash); err != nil {
		return err
	}
	p.hash = hash
	return nil
}

// GetPlain is the plain getter
func (p *Hash) GetPlain() string {
	return p.plain
}

// SetPlain is the plain setter
func (p *Hash) SetPlain(plain string) {
	p.plain = plain
}

// Compare will matches hashes using the plain password
func (p *Hash) Compare(plain string) bool {
	return strings.Compare(p.hasher.GetHash(plain), p.hash) == 0
}
