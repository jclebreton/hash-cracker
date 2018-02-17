package cmd

import (
	"github.com/jclebreton/hash-cracker/hashers"
	"github.com/pkg/errors"
)

// Hash is the struct to store data
type Hash struct {
	hasher hashers.Hasher
	hash   string
	plain  string
}

// SetHasher is the hasher setter
func (h *Hash) SetHasher(hasher hashers.Hasher) {
	h.hasher = hasher
}

// GetHash is the hash getter
func (h *Hash) GetHash() string {
	return h.hash
}

// SetHash is the hash and salt setter
func (h *Hash) SetHash(hash string) error {
	h.hash = hash
	return nil
}

// GetPlain is the plain getter
func (h *Hash) GetPlain() string {
	return h.plain
}

// SetPlain is the plain setter
func (h *Hash) SetPlain(plain string) {
	h.plain = plain
}

// Compare will matches hashes using the plain password
func (h *Hash) Compare(plain string) (bool, error) {
	hash, err := h.hasher.Compare(h.hash, plain)
	if err != nil {
		return false, errors.Wrap(err, "Compare error")
	}
	return hash == h.hash, nil
}
