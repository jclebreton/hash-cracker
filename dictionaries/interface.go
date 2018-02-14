package dictionaries

// Provider is the interface used to provide dictionaries or hashes
type Provider interface {
	GetName() string

	Prepare() error
	Next() bool
	Value() string
	Err() error
	Close() error

	GetTotal() int
}
