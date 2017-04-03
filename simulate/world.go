package simulate

type World struct {
	Warehouse Warehouse

	Mine Mine
	Lumberyard Lumberyard
	Workshop Workshop

	Workers []*Worker
}

func (w *World) Run() {
	for _, worker := range w.Workers {
		go worker.run()
	}
}
