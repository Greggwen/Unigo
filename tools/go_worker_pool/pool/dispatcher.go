package pool

import "log"

var WorkerChannel = make(chan chan Work)

type Collector struct {
	Work chan Work
	End  chan bool
}

func StartDispatcher(workerCount int) Collector {
	var i int
	var workers []Worker

	input := make(chan Work)
	end := make(chan bool)

	collector := Collector{input, end}

	for i < workerCount {
		i++
		log.Println("Starting Worker ", i)
		worker := Worker{
			ID:            i,
			Channel:       make(chan Work),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool),
		}

		worker.Start()
		workers = append(workers, worker)

	}

	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					w.Stop()
				}
				return
			case work := <-input:
				worker := <-WorkerChannel
				worker <- work
			}
		}
	}()

	return collector
}
