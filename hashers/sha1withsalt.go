package hashers

import (
	"crypto/sha1"
	"fmt"

	"github.com/pkg/errors"
)

// Sha1WithSalt use the hash mode: $salt.sha1($salt.$pass)
type Sha1WithSalt struct {
	salt string
}

// SetSaltFromHash is the salt setter
func (l *Sha1WithSalt) SetSaltFromHash(hash string) error {
	if len(hash) != 56 {
		return errors.New("not a valid hash")
	}
	l.salt = hash[0:16]
	return nil
}

// GetHash is the hash getter
func (l *Sha1WithSalt) GetHash(plain string) string {
	h := sha1.New()
	h.Write([]byte(l.salt + plain))
	return fmt.Sprintf("%s%x", l.salt, h.Sum(nil))
}
