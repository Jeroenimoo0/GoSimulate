package simulate

import (
	"time"
)

type Worker struct {
	Storage
	name string
	world *World
	workplace *Workplace
	task Task
}

func NewWorker(name string, world *World) *Worker {
	return &Worker{
		*NewStorage(3),
		name,
		world,
		nil,
		nil,
	}
}

func (w *Worker) step() time.Duration {
	if w.task != nil {
		w.task.Complete(w)
		w.task = nil
	}

	if w.workplace == nil {
		return time.Second * -1
	} else {
		task := w.workplace.fetchWork(w)
		if task != nil {
			w.task = task
			w.task.Start(w)
			return w.task.GetTimeToComplete(w)
		} else {
			w.task = nil
			return time.Second
		}
	}
}

func (w *Worker) run() {
	for {
		sleep := w.step()
		if sleep < 0 {
			return
		}

		time.Sleep(sleep)
	}
}
