package main

import (
	"fmt"
	"github.com/Jeroenimoo0/GoSimulate/simulate"
	"time"
)

import (
	_ "net/http/pprof"
	"log"
	"net/http"
)

func dataSmall() {
	world := simulate.World{}
	world.Warehouse = *simulate.NewWarehouse(&world)
	world.Mine = *simulate.NewMine(&world)
	world.Lumberyard = *simulate.NewLumberyard(&world)
	world.Workshop = *simulate.NewWorkshop(&world)
	world.Workers = make([]*simulate.Worker, 0, 3)

	world.Workers = append(world.Workers, simulate.NewWorker("Bob", &world))
	world.Workers = append(world.Workers, simulate.NewWorker("Bobby", &world))

	for i := 0; i < 6; i++ {
		worker := simulate.NewWorker("Harry", &world)
		world.Workers = append(world.Workers, worker)

		if !world.Mine.Workplace.Full() {
			world.Mine.Assign(worker)
		} else if !world.Lumberyard.Workplace.Full() {
			world.Lumberyard.Assign(worker)
		} else if !world.Workshop.Workplace.Full() {
			world.Workshop.Assign(worker)
		}
	}

	for i := 0; i < 200; i++ {
		world.Workshop.AddTask(simulate.NewTaskResource(
			"Craft axe",
			&world.Warehouse.Storage,
			time.Second * 4,
			[]simulate.Resource{simulate.Stone, simulate.Wood, simulate.Wood},
			[]simulate.Resource{},
		))
	}

	world.Run()
}

// ----------- Setup & run ------------

func main() {
	fmt.Println("Hello world!")

	for i := 0; i < 10000; i++ {
		dataSmall()
		fmt.Println("Created world:", i)
	}

	log.Println(http.ListenAndServe("localhost:6060", nil))

	quit := make(chan bool)
	for {
		<- quit
	}

	fmt.Println("Bye world.")
}