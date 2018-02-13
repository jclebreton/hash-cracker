package comparators

import (
	"crypto/sha1"
	"fmt"
)

type LbcHash struct {
	salt string
}

func (l *LbcHash) setSalt(salt string) {
	l.salt = salt
}

func (l *LbcHash) getHash(plain string) string {
	h := sha1.New()
	h.Write([]byte(l.salt + plain))
	return fmt.Sprintf("%s%x", l.salt, h.Sum(nil))
}
