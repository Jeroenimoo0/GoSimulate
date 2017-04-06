package main


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

func (w *Worker) Run() {
	if w.task != nil {
		//fmt.Println(w.name, "completed", w.task.GetName())
		w.task.Complete(w)
		w.task = nil
	}

	if w.workplace == nil {
		return
	} else {
		task := w.workplace.fetchWork(w)
		if task != nil {
			w.task = task
			w.task.Start(w)

			//fmt.Println(w.name, "started", w.task.GetName())

			w.world.Simulation.Add(w, w.task.GetTimeToComplete(w))
		} else {
			//TODO Add callback for job
			w.task = nil
		}
	}
}
