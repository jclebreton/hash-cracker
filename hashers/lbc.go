package hashers

import (
	"crypto/sha1"
	"fmt"

	"github.com/pkg/errors"
)

// LbcHash use the hash mode: sha1($salt.$pass)
type LbcHash struct {
	salt string
}

// SetSaltFromHash is the salt setter
func (l *LbcHash) SetSaltFromHash(hash string) error {
	if len(hash) != 56 {
		return errors.New("not a valid lbc hash")
	}
	l.salt = hash[0:16]
	return nil
}

// GetHash is the hash getter
func (l *LbcHash) GetHash(plain string) string {
	h := sha1.New()
	h.Write([]byte(l.salt + plain))
	return fmt.Sprintf("%s%x", l.salt, h.Sum(nil))
}
