package hashers

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
)

// Sha1WithSalt use the hash mode: $salt.sha1($salt.$pass)
type Sha1WithSalt struct {
}

func (s *Sha1WithSalt) getSaltFromHash(hash string) (string, error) {
	if len(hash) != 56 {
		return "", errors.New("unable to get salt from hash")
	}
	return hash[0:16], nil
}

// GetHash is the hash getter
func (s *Sha1WithSalt) Compare(hash, plain string) (string, error) {
	salt, err := s.getSaltFromHash(hash)
	if err != nil {
		return "", errors.Wrap(err, "Compare error")
	}

	h := sha1.New()
	_, err = h.Write([]byte(salt + plain))
	if err != nil {
		return "", errors.Wrap(err, "Compare error")
	}

	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("%s%s", salt, sha1_hash), nil
}
