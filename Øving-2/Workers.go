package main

import "sync"

// Workers : list of worker_threads
type Workers struct {
	sync.Mutex
	cond  *sync.Cond
	count int
}

// NewWorkers : creates a new Workers object
func NewWorkers(count int) *Workers {
	w := Workers{count: count}
	w.cond = sync.NewCond(&w)
	return &w
}

// Start : starts worker thread(s)
func (w Workers) Start() {
	for i := 0; i < w.count; i++ {

	}
}

func (w Workers) Post() {

}
