package hashers

// Hasher is the interface used to hash plain passwords
type Hasher interface {
	Compare(hash, plain string) (string, error)
}
