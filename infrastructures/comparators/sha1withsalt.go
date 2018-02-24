package comparators

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/jclebreton/hash-cracker/domains"
	"github.com/pkg/errors"
)

// Sha1WithSalt use the hash mode: $salt.sha1($salt.$pass)
type Sha1WithSalt struct {
}

// Compare
func (comparator *Sha1WithSalt) Compare(hash domains.Hash, plain string) (bool, error) {
	// Retrieve salt
	if len(hash.Hash) != 56 {
		return false, errors.New("unable to get salt from hash")
	}
	salt := hash.Hash[0:16]

	// Build new hash
	h := sha1.New()
	h.Write([]byte(salt + plain)) // no errors
	newHash := hex.EncodeToString(h.Sum(nil))
	newHash = fmt.Sprintf("%s%s", salt, newHash)

	return hash.Hash == newHash, nil
}
