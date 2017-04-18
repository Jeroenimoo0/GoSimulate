package main

import (
	"time"
	"testing"
)

type fCreateWorld func() *World

func BenchmarkMain_SimulateWorldSmall4(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 4, b)}
func BenchmarkMain_SimulateWorldSmall8(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 8, b)}
func BenchmarkMain_SimulateWorldSmall12(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 12, b)}
func BenchmarkMain_SimulateWorldSmall16(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 16, b)}
func BenchmarkMain_SimulateWorldSmall20(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 20, b)}
func BenchmarkMain_SimulateWorldSmall24(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 24, b)}
func BenchmarkMain_SimulateWorldSmall28(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 28, b)}
func BenchmarkMain_SimulateWorldSmall32h(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 32, b)}
func BenchmarkMain_SimulateWorldSmall36h(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 36, b)}
func BenchmarkMain_SimulateWorldSmall40h(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 40, b)}
func BenchmarkMain_SimulateWorldSmall44h(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 44, b)}
func BenchmarkMain_SimulateWorldSmall48h(b *testing.B) { benchmarkMain_SimulateWorld(smallWorld, time.Hour * 48, b)}

func BenchmarkMain_SimulateWorldMedium4(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 4, b)}
func BenchmarkMain_SimulateWorldMedium8(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 8, b)}
func BenchmarkMain_SimulateWorldMedium12(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 12, b)}
func BenchmarkMain_SimulateWorldMedium16(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 16, b)}
func BenchmarkMain_SimulateWorldMedium20(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 20, b)}
func BenchmarkMain_SimulateWorldMedium24(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 24, b)}
func BenchmarkMain_SimulateWorldMedium28(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 28, b)}
func BenchmarkMain_SimulateWorldMedium32h(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 32, b)}
func BenchmarkMain_SimulateWorldMedium36h(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 36, b)}
func BenchmarkMain_SimulateWorldMedium40h(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 40, b)}
func BenchmarkMain_SimulateWorldMedium44h(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 44, b)}
func BenchmarkMain_SimulateWorldMedium48h(b *testing.B) { benchmarkMain_SimulateWorld(mediumWorld, time.Hour * 48, b)}

func BenchmarkMain_SimulateWorldBig4(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 4, b)}
func BenchmarkMain_SimulateWorldBig8(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 8, b)}
func BenchmarkMain_SimulateWorldBig12(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 12, b)}
func BenchmarkMain_SimulateWorldBig16(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 16, b)}
func BenchmarkMain_SimulateWorldBig20(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 20, b)}
func BenchmarkMain_SimulateWorldBig24(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 24, b)}
func BenchmarkMain_SimulateWorldBig28(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 28, b)}
func BenchmarkMain_SimulateWorldBig32h(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 32, b)}
func BenchmarkMain_SimulateWorldBig36h(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 36, b)}
func BenchmarkMain_SimulateWorldBig40h(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 40, b)}
func BenchmarkMain_SimulateWorldBig44h(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 44, b)}
func BenchmarkMain_SimulateWorldBig48h(b *testing.B) { benchmarkMain_SimulateWorld(bigWorld, time.Hour * 48, b)}

func benchmarkMain_SimulateWorld(f fCreateWorld, l time.Duration, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		world := f()

		b.StartTimer()

		world.Simulation.Run()
	}
}
