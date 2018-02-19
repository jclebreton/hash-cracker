package progress

import "gopkg.in/cheggaaa/pb.v1"

type CheggaaaPool struct {
	pool *pb.Pool
	bars []*pb.ProgressBar
}

func (p *CheggaaaPool) Add(bar ProgressBarer) {
	p.bars = append(p.bars, bar.GetBar().Bar)
}

func (p *CheggaaaPool) Start() error {
	pool, err := pb.StartPool(p.bars...)
	if err != nil {
		return err
	}
	p.pool = pool
	return nil
}

func (p *CheggaaaPool) Stop() error {
	return p.pool.Stop()
}
