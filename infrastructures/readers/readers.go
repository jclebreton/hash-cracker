package readers

// DictionaryReader is the interface used to provide dictionaries
type DictionaryProvider interface {
	Prepare() error
	Next() bool
	Value() string
	Err() error
	Close() error

	GetName() string
	GetTotal() int64
}

// HashesReader is the interface used to provide hashes
type HashesProvider interface {
	Prepare() error
	Next() bool
	Value() string
	Err() error
	Close() error

	GetName() string
	GetTotal() int64
}
