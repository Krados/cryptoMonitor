package lib

type Job interface {
	Exec()
}

type Dispatcher struct {
	Workers []*Worker
	JobChan chan Job
}

func NewDispatcher(workerNum int) Dispatcher {
	d := Dispatcher{
		Workers: make([]*Worker, workerNum),
		JobChan: make(chan Job),
	}
	for i := 0; i < workerNum; i++ {
		d.Workers[i] = &Worker{ID: i}
	}

	return d
}

func (d *Dispatcher) Start() {
	for _, worker := range d.Workers {
		go worker.Listen(d.JobChan)
	}
}

func (d *Dispatcher) Dispatch(job Job) {
	d.JobChan <- job
}

type Worker struct {
	ID int
}

func (w *Worker) Listen(jobChan chan Job) {
	for j := range jobChan {
		j.Exec()
	}
}
