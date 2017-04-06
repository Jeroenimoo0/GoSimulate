package main

import (
	"fmt"
	"time"
	_ "net/http/pprof"
	"github.com/Jeroenimoo0/GoSimulate"
)

func smallWorld() *World {
	return world(1)
}

func mediumWorld() *World {
	return world(3)
}

func bigWorld() *World {
	return world(5)
}

func world(multiplier int) *World {
	world := World{}
	world.Simulation = *simulate.NewSimulation()

	world.Warehouse = *NewWarehouse(&world)
	world.Mine = []Mine{}
	world.Lumberyard = []Lumberyard{}
	world.Workshop = []Workshop{}
	world.Workers = []*Worker{}

	for i := 0; i < multiplier; i++ {
		world.Mine = append(world.Mine, *NewMine(&world))
		world.Lumberyard = append(world.Lumberyard, *NewLumberyard(&world))
		world.Workshop = append(world.Workshop, *NewWorkshop(&world))

		world.Workshop[i].workSupplier = TaskSupplier{
			NewTaskResource(
				"Craft axe",
				&world.Warehouse.Storage,
				time.Second * 4,
				[]Resource{Stone, Wood, Wood},
				[]Resource{},
			),
		}
	}

	for i := 0; i < 6 * multiplier; i++ {
		worker := NewWorker(fmt.Sprint("Worker", i), &world)
		world.Workers = append(world.Workers, worker)

		for j := 0; j < multiplier; j++ {
			if !world.Mine[j].Workplace.Full() {
				world.Mine[j].Assign(worker)
				break
			} else if !world.Lumberyard[j].Workplace.Full() {
				world.Lumberyard[j].Assign(worker)
				break
			} else if !world.Workshop[j].Workplace.Full() {
				world.Workshop[j].Assign(worker)
				break
			}
		}

		world.Simulation.Add(worker, 0)
	}

	return &world
}


// ----------- Setup & run ------------

func main() {
	for i := 0; i < 1; i++ {
		world := smallWorld()

		now := time.Now()
		world.Simulation.Run(time.Second * 60 * 60 * 4)

		fmt.Println("Simulated world for ", world.Simulation.TotalTime, "in", time.Now().Sub(now))
	}

	//log.Println(http.ListenAndServe("localhost:6060", nil))

	fmt.Println("Bye world.")
}