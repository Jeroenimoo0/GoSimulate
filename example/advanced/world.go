package main

import "github.com/Jeroenimoo0/GoSimulate"

type World struct {
	Simulation simulate.Simulation

	Warehouse Warehouse

	Mine []Mine
	Lumberyard []Lumberyard
	Workshop []Workshop

	Workers []*Worker
}
