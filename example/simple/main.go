package main

import (
	"fmt"
	"github.com/Jeroenimoo0/GoSimulate"
	"time"
)

var simulation simulate.Simulation = simulate.NewSimulationInstant(time.Second * 20)

type Worker struct {
	name string
	saidHello bool
}

func (w *Worker) Run() {

	if w.saidHello {
		fmt.Println(w.name, "says goodbye")
	} else {
		w.saidHello = true
		fmt.Println(w.name, "says hello!")
		simulation.Add(w, time.Second)
	}
}

func main() {
	simulation.Add(&Worker{"Lolo", false}, time.Second * 5)
	simulation.Add(&Worker{"John", false}, time.Millisecond * 6320)
	simulation.Add(&Worker{"Bob", false}, time.Millisecond * 4140)
	simulation.Add(&Worker{"Harry", false}, time.Second)

	simulation.Run()
}
