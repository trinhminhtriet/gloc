package bufferpool

import (
	"bytes"
	"sync"
)

type Pool struct {
	sync.Pool
}

func NewPool() *Pool {
	return &Pool{
		sync.Pool{New: func() any { return new(bytes.Buffer) }},
	}
}

func (p *Pool) Get() *bytes.Buffer {
	buf := p.Pool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func (p *Pool) Put(buf *bytes.Buffer) {
	p.Pool.Put(buf)
}
