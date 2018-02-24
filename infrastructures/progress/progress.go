package progress

// Barer is the interface used to manipulate a progress Bar
type Barer interface {
	Set(value int)
	SetTotal(total int64)
	Add(add int64)
	Increment() int64
	IncrementTotal(add int64)
	GetBar() *CheggaaBar
	Start()
	Finish()
}

// BarPooler is the interface to manipulate a lot of progress bars
type BarPooler interface {
	Add(bar Barer)
	Start() error
	Stop() error
}
