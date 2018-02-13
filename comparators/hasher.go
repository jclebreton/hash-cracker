package comparators

import "strings"

// Hasher is the interface used to hash plain passwords
type Hasher interface {
	setSalt(salt string)
	getHash(plain string) string
}

// Hash is the struct to store LBC authentication fields
type Hash struct {
	hasher Hasher
	hash   string
	plain  string
}

// SetSalt is the salt setter
func (p *Hash) SetHasher(hasher Hasher) {
	p.hasher = hasher
}

// SetSalt is the salt setter
func (p *Hash) SetSalt(salt string) {
	p.hasher.setSalt(salt)
}

// SetHash is the hash setter
func (p *Hash) SetHash(hash string) {
	p.hash = hash
}

// GetHash is the hash getter
func (p *Hash) GetHash() string {
	return p.hash
}

// GetPlain is the plain getter
func (p *Hash) GetPlain() string {
	return p.plain
}

// Compare will matches hashes using the plain password
func (p *Hash) Compare(plain string) bool {
	return strings.Compare(p.hasher.getHash(plain), p.hash) == 0
}
