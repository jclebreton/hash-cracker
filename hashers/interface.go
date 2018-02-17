package hashers

// Hasher is the interface used to hash plain passwords
type Hasher interface {
	SetSaltFromHash(hash string) error
	GetHash(plain string) (string, error)
}
