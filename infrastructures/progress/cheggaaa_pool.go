package progress

import "gopkg.in/cheggaaa/pb.v1"

// CheggaaaPool is a progress Bar pool implementation
type CheggaaaPool struct {
	pool *pb.Pool
	bars []*pb.ProgressBar
}

// Add registers a new progress bar
func (p *CheggaaaPool) Add(bar Barer) {
	p.bars = append(p.bars, bar.GetBar().Bar)
}

// Start launch all registered progress bars
func (p *CheggaaaPool) Start() error {
	pool, err := pb.StartPool(p.bars...)
	if err != nil {
		return err
	}
	p.pool = pool
	return nil
}

// Stop stops the pool
func (p *CheggaaaPool) Stop() error {
	return p.pool.Stop()
}
