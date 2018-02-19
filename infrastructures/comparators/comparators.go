package comparators

import "github.com/jclebreton/hash-cracker/domains"

// Comparator is the interface used compare hash and plain password
type Comparator interface {
	Compare(hash domains.Hash, plain string) (bool, error)
}
