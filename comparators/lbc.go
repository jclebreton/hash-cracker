package comparators

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// LBCPassword is the struct to store LBC authentication fields
type LBCPassword struct {
	Hash  string
	Salt  string
	Plain string
}

// GetHash returns the current LBC hash
func (p *LBCPassword) GetHash() string {
	return p.Hash
}

// SetHash initializes current struct from LBC hash
func (p *LBCPassword) SetHash(hash string) {
	p.Salt = hash[0:16]
	p.Hash = hash
	p.Plain = ""
}

// Compare will matches plain password with LBC hash
func (p *LBCPassword) Compare(plain string) bool {
	h := sha1.New()
	h.Write([]byte(p.Salt + plain))
	hash := p.Salt + fmt.Sprintf("%x", h.Sum(nil))
	return strings.Compare(hash, p.Hash) == 0
}
