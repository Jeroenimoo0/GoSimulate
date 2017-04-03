package simulate

import (
	"time"
	"sync"
)

type Building struct {
	world *World
}

type WorkSupplier interface {
	findWork(worker *Worker) Task
}

type Workplace struct {
	workSpots []*Worker
	tL sync.RWMutex
	tasks []Task
	workSupplier WorkSupplier
}

func NewWorkplace(spots int) *Workplace {
	return &Workplace{
		make([]*Worker, 0, spots),
		sync.RWMutex{},
		make([]Task, 0, 5),
		nil,
	}
}

func (w *Workplace) findWork() Task {
	return nil
}

func (w *Workplace) fetchWork(worker *Worker) Task {
	var task Task = nil

	// First try to find a task in queue
	w.tL.Lock()
	if len(w.tasks) > 0 {
		for i, task := range w.tasks {
			if task.CanBeCompleted(worker) {
				w.tasks = append(w.tasks[:i], w.tasks[i+1:]...)

				w.tL.Unlock()
				return task
			}
		}
	}
	w.tL.Unlock()

	// If not possible try to see if the workplace offers
	// default tasks
	if w.workSupplier != nil {
		task = w.workSupplier.findWork(worker)
	}

	if task != nil && task.CanBeCompleted(worker) {
		return task
	}

	return nil
}

func (w *Workplace) Assign(worker *Worker) {
	if w.Full() {
		panic("Workspots were already full!")
	}

	w.workSpots = append(w.workSpots, worker)
	worker.workplace = w
}

func (w *Workplace) Full() bool {
	return len(w.workSpots) == cap(w.workSpots)
}

func (w *Workplace) AddTask(task Task) {
	w.tL.Lock()

	w.tasks = append(w.tasks, task)

	w.tL.Unlock()
}

// Buildings ----------------------------

type Lumberyard struct {
	Building
	Workplace
	Storage
}

func NewLumberyard(world *World) *Lumberyard {
	lumberyard := Lumberyard{
		Building{world},
		*NewWorkplace(2),
		*NewStorage(10),
	}

	lumberyard.workSupplier = &lumberyard

	return &lumberyard
}

func (l *Lumberyard) findWork(worker *Worker) Task {
	if l.Storage.Full() {
		if l.world.Warehouse.Full() {
			return nil
		} else {
			var task Task = NewTaskMoving(
				"transfer wood",
				time.Second * 2,
				Wood,
				&l.Storage,
				&l.world.Warehouse.Storage,
			)

			return task
		}
	} else {
		if l.Storage.Full() {
			return nil
		} else {
			var task Task = NewTaskResource(
				"Cut tree",
				&l.Storage,
				time.Second * 2,
				[]Resource{},
				[]Resource{Wood},
			)

			return task
		}
	}

	return nil
}

type Mine struct {
	Building
	Workplace
	Storage
}

func NewMine(world *World) *Mine {
	mine := Mine{
		Building{world},
		*NewWorkplace(2),
		*NewStorage(10),
	}

	mine.workSupplier = &mine

	return &mine
}

func (m *Mine) findWork(worker *Worker) Task {
	if m.Storage.Full() {
		if m.world.Warehouse.Full() {
			return nil
		} else {
			var task Task = NewTaskMoving(
				"transfer stone",
				time.Second * 2,
				Stone,
				&m.Storage,
				&m.world.Warehouse.Storage,
			)

			return task
		}
	} else {
		if m.Storage.Full() {
			return nil
		} else {
			var task Task = NewTaskResource(
				"Mine stone",
				&m.Storage,
				time.Second * 2,
				[]Resource{},
				[]Resource{Stone},
			)

			return task
		}
	}

	return nil
}

type Workshop struct {
	Building
	Workplace
	Storage
}

func NewWorkshop(world *World) *Workshop {
	return &Workshop{
		Building{world},
		*NewWorkplace(2),
		*NewStorage(10),
	}
}

type Warehouse struct {
	Building
	Storage
}

func NewWarehouse(world *World) *Warehouse {
	return &Warehouse{
		Building{world},
		*NewStorage(100),
	}
}
