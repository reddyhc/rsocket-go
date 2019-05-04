package transport

import "sync/atomic"

type Counter struct {
	r, w uint64
}

func (p *Counter) incrWriteBytes(n int) {
	atomic.AddUint64(&(p.w), uint64(n))
}

func (p *Counter) incrReadBytes(n int) {
	atomic.AddUint64(&(p.r), uint64(n))
}

func (p *Counter) ReadBytes() uint64 {
	return atomic.LoadUint64(&(p.r))
}

func (p *Counter) WriteBytes() uint64 {
	return atomic.LoadUint64(&(p.w))
}

func NewCounter() *Counter {
	return &Counter{}
}
