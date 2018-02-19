package progress

// CheggaaBar is the interface used to manipulate a progress Bar
type ProgressBarer interface {
	Set(value int)
	SetTotal(total int64)
	Add(add int64)
	Increment() int64
	IncrementTotal(add int64)
	GetBar() *CheggaaBar
	Start()
	Finish()
}

// ProgressBarPooler is the interface to manipulate a lot of progress bars
type ProgressBarPooler interface {
	Add(bar ProgressBarer)
	Start() error
	Stop() error
}
