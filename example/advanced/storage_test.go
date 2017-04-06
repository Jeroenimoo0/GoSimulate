package main

import (
	"testing"
	"time"
)

func TestEqualEqual1(t *testing.T) {
	first := []Resource{None, Wood, None}
	second := []Resource{None, Wood, None}

	if Equal(first, second) != true {
		t.Fail()
	}
}

func TestEqualNotEqual1(t *testing.T) {
	first := []Resource{None, Wood, None}
	second := []Resource{None, None, None}

	if Equal(first, second) != false {
		t.Fail()
	}
}

func TestEqualNotEqual2(t *testing.T) {
	first := []Resource{None, Wood, None}
	second := []Resource{None, Stone, None}

	if Equal(first, second) != false {
		t.Fail()
	}
}

func TestEqualNotEqual3(t *testing.T) {
	first := []Resource{None, None, None, None}
	second := []Resource{None, None, None}

	if Equal(first, second) != false {
		t.Fail()
	}
}

func TestStorage_TransferOrWait(t *testing.T) {
	s1 := NewStorage(10)
	s2 := NewStorage(10)

	go func() {
		time.Sleep(1)
		s1.Add(Wood)
		time.Sleep(1)
		s1.Add(Wood)
	}()

	s1.TransferOrWait(s2, Wood)
	s1.TransferOrWait(s2, Wood)

	if !s2.Contains(Wood) || s1.Contains(Wood) {
		t.Fail()
	}
}

func benchmarkStorage_GetResources(size int, b *testing.B) {
	b.StopTimer()
	var storage Storage = *NewStorage(size)
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		storage.GetResources()
	}
}

func BenchmarkStorage_GetResources3(b *testing.B) { benchmarkStorage_GetResources(3, b)}
func BenchmarkStorage_GetResources10(b *testing.B) { benchmarkStorage_GetResources(10, b)}
func BenchmarkStorage_GetResources20(b *testing.B) { benchmarkStorage_GetResources(20, b)}
func BenchmarkStorage_GetResources30(b *testing.B) { benchmarkStorage_GetResources(30, b)}
func BenchmarkStorage_GetResources50(b *testing.B) { benchmarkStorage_GetResources(50, b)}
func BenchmarkStorage_GetResources100(b *testing.B) { benchmarkStorage_GetResources(100, b)}

func benchmarkStorage_Full(size int, fill bool, b *testing.B) {
	b.StopTimer()
	var storage Storage = *NewStorage(size)

	if fill {
		for i := 0; i < size; i++ {
			storage.Add(Wood)
		}
	}

	b.StartTimer()

	for n := 0; n < b.N; n++ {
		storage.Full()
	}
}

func BenchmarkStorage_Full3Full(b *testing.B) { benchmarkStorage_Full(3, true, b)}
func BenchmarkStorage_Full10Full(b *testing.B) { benchmarkStorage_Full(10, true, b)}
func BenchmarkStorage_Full20Full(b *testing.B) { benchmarkStorage_Full(20, true, b)}
func BenchmarkStorage_Full30Full(b *testing.B) { benchmarkStorage_Full(30, true, b)}
func BenchmarkStorage_Full50Full(b *testing.B) { benchmarkStorage_Full(50, true, b)}
func BenchmarkStorage_Full100Full(b *testing.B) { benchmarkStorage_Full(100, true, b)}

func BenchmarkStorage_Full3Empty(b *testing.B) { benchmarkStorage_Full(3, false, b)}
func BenchmarkStorage_Full10Empty(b *testing.B) { benchmarkStorage_Full(10, false, b)}
func BenchmarkStorage_Full20Empty(b *testing.B) { benchmarkStorage_Full(20, false, b)}
func BenchmarkStorage_Full30Empty(b *testing.B) { benchmarkStorage_Full(30, false, b)}
func BenchmarkStorage_Full50Empty(b *testing.B) { benchmarkStorage_Full(50, false, b)}
func BenchmarkStorage_Full100Empty(b *testing.B) { benchmarkStorage_Full(100, false, b)}
