package progress

import "gopkg.in/cheggaaa/pb.v1"

type CheggaaBar struct {
	Bar *pb.ProgressBar
}

func NewProgressBar(title string) *CheggaaBar {
	bar := pb.New(0).SetUnits(pb.U_NO).Prefix(title)
	bar.ShowPercent = true
	bar.ShowTimeLeft = true
	return &CheggaaBar{bar}
}

func (pb *CheggaaBar) GetBar() *CheggaaBar {
	return pb
}

func (pb *CheggaaBar) SetTotal(total int64) {
	pb.Bar.Total = total
	pb.Bar.ShowPercent = true
	pb.Bar.ShowTimeLeft = true

}

func (pb *CheggaaBar) IncrementTotal(add int64) {
	pb.Bar.Total += add
}

func (pb *CheggaaBar) Add(add int64) {
	pb.Bar.Add(int(add))
}

func (pb *CheggaaBar) Increment() int64 {
	return int64(pb.Bar.Increment())
}

func (pb *CheggaaBar) Start() {
	pb.Bar.Start()
}

func (pb *CheggaaBar) Finish() {
	pb.Bar.Finish()
}

func (pb *CheggaaBar) Set(value int) {
	pb.Bar.Set(value)
}
