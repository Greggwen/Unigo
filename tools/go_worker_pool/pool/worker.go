package pool

import (
	"Unigo/tools/go_worker_pool/work"
	"log"
)

// Work is Struct
type Work struct {
	ID  int
	Job string
}

// Worker is a Struct
type Worker struct {
	ID            int
	WorkerChannel chan chan Work
	Channel       chan Work
	End           chan bool
}

// Start func is Start
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case job := <-w.Channel:
				work.DoWork(job.Job, job.ID)
			case <-w.End:
				return
			}
		}
	}()
}

// Stop is a func
func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping", w.ID)
	w.End <- true
}
