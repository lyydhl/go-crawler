package scheduler

import (
	"go-crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (q *SimpleScheduler) WorkerReady(w chan engine.Request) {
}

func (q *SimpleScheduler) WorkerChan() chan engine.Request {
	return q.workerChan
}

func (q *SimpleScheduler) Run() {
}
