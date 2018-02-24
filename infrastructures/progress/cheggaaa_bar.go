package progress

import "gopkg.in/cheggaaa/pb.v1"

// CheggaaBar is a progress Bar implementation
type CheggaaBar struct {
	Bar *pb.ProgressBar
}

// NewProgressBar is the constructor
func NewProgressBar(title string) *CheggaaBar {
	bar := pb.New(0).SetUnits(pb.U_NO).Prefix(title)
	bar.ShowPercent = true
	bar.ShowTimeLeft = true
	return &CheggaaBar{bar}
}

// GetBar returns all the structure
func (pb *CheggaaBar) GetBar() *CheggaaBar {
	return pb
}

// SetTotal is the total setter
func (pb *CheggaaBar) SetTotal(total int64) {
	pb.Bar.Total = total
	pb.Bar.ShowPercent = true
	pb.Bar.ShowTimeLeft = true

}

// IncrementTotal allows to increase the total progress bar
func (pb *CheggaaBar) IncrementTotal(add int64) {
	pb.Bar.Total += add
}

// Add allows to increase the current progress
func (pb *CheggaaBar) Add(add int64) {
	pb.Bar.Add(int(add))
}

// Increment allows to increment the current progress
func (pb *CheggaaBar) Increment() int64 {
	return int64(pb.Bar.Increment())
}

// Start starts the progress bar
func (pb *CheggaaBar) Start() {
	pb.Bar.Start()
}

// Finish stops the progress bar
func (pb *CheggaaBar) Finish() {
	pb.Bar.Finish()
}

// Set allows to reset the current progress of the progress bar
func (pb *CheggaaBar) Set(value int) {
	pb.Bar.Set(value)
}
