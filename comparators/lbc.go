package comparators

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var r *rand.Rand // Rand for this package.

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomString(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := ""
	for i := 0; i < strlen; i++ {
		index := r.Intn(len(chars))
		result += chars[index : index+1]
	}
	return result
}

func hashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// LBCPassword is the struct to store LBC authentication fields
type LBCPassword struct {
	Hash  string
	Salt  string
	Plain string
}

// SetRandomSalt generates a random salt and reset other fields
func (p *LBCPassword) SetRandomSalt() {
	p.Salt = randomString(16)
	p.Plain = ""
	p.Hash = ""
}

// SetPlainPassword initializes current struct from a plain password
func (p *LBCPassword) SetPlainPassword(password string) error {
	if len(p.Salt) != 16 {
		return errors.New("salt error")
	}
	p.Hash = p.Salt + hashString(p.Salt+password)
	p.Plain = password
	return nil
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
	hash := p.Salt + hashString(p.Salt+plain)
	return strings.Compare(hash, p.Hash) == 0
}
