package lib

import "errors"

type Work interface {
	Action()
}

type WorkRunner struct {
	WorkChan chan Work
}

func NewWorkRunner(workLimit int) *WorkRunner {
	return &WorkRunner{
		WorkChan: make(chan Work, workLimit),
	}
}

func (w *WorkRunner) Receive(work Work) (err error) {
	select {
	case w.WorkChan <- work:
		go w.run(work)
	default:
		err = errors.New("work list is full")
	}
	return
}

func (w *WorkRunner) run(work Work) {
	work.Action()
	w.pop()
}

func (w *WorkRunner) pop() {
	<-w.WorkChan
}
