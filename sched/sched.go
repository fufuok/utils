// Copyright 2022 Changkun Ou <changkun.de>. All rights reserved.
// Use of this source code is governed by a GPLv3 license that
// can be found in the LICENSE file.

package sched

import (
	"runtime"
	"sync/atomic"
)

// Pool is a worker pool.
type Pool struct {
	running    uint64
	numWorkers int
	numQueues  int
	tasks      chan funcdata
	done       chan struct{}
}

type funcdata struct {
	fn func()
	fg func(interface{})
	ar interface{}
}

// Option is a scheduler option.
type Option func(w *Pool)

// Workers is number of workers that can execute tasks concurrently.
func Workers(limit int) Option {
	return func(w *Pool) {
		if limit > 0 {
			w.numWorkers = limit
		}
	}
}

// Queues is buffer capacity of the tasks channel.
func Queues(limit int) Option {
	return func(w *Pool) {
		if limit >= 0 {
			w.numQueues = limit
		}
	}
}

// New creates a new task scheduler and returns a pool of workers.
func New(opts ...Option) *Pool {
	n := runtime.NumCPU()
	p := &Pool{
		running:    0,
		numWorkers: n,
		numQueues:  n * 100,
		done:       make(chan struct{}),
	}

	for _, opt := range opts {
		opt(p)
	}

	p.tasks = make(chan funcdata, p.numQueues)

	// Start workers
	for i := 0; i < p.numWorkers; i++ {
		go func() {
			for d := range p.tasks {
				if d.fn != nil {
					d.fn()
				} else {
					d.fg(d.ar)
				}
				p.complete()
			}
		}()
	}

	return p
}

// Run runs f in the current pool.
func (p *Pool) Run(f ...func()) {
	for i := range f {
		p.tasks <- funcdata{fn: f[i]}
	}
}

func (p *Pool) RunWithArgs(f func(args interface{}), args interface{}) {
	p.tasks <- funcdata{fg: f, ar: args}
}

func (p *Pool) Add(numTasks int) int {
	return int(atomic.AddUint64(&p.running, uint64(numTasks)))
}

func (p *Pool) Running() uint64 {
	return atomic.LoadUint64(&p.running)
}

func (p *Pool) IsRunning() bool {
	return p.Running() != 0
}

func (p *Pool) Wait() {
	<-p.done
}

func (p *Pool) Release() {
	close(p.tasks)
	close(p.done)
}

func (p *Pool) WaitAndRelease() {
	p.Wait()
	p.Release()
}

func (p *Pool) complete() {
	ret := atomic.AddUint64(&p.running, ^uint64(0))
	if ret == 0 {
		p.done <- struct{}{}
	}
}
