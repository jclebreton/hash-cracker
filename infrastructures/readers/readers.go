package readers

// DictionaryProvider is the interface used to provide dictionaries
type DictionaryProvider interface {
	Prepare() error
	Next() bool
	Value() string
	Err() error
	Close() error

	GetName() string
	GetTotal() int64
}

// HashesProvider is the interface used to provide hashes
type HashesProvider interface {
	Prepare() error
	Next() bool
	Value() string
	Err() error
	Close() error

	GetName() string
	GetTotal() int64
}
